package main

import (
	"context"
	"errors"
	"fmt"
	"goth/internal/auth/tokenauth"
	"goth/internal/handlers"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	m "goth/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"

	"database/sql"
	postgres "goth/internal/db"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func getDBConnectionString() string {
	var (
		db_database = os.Getenv("DB_DATABASE")
		db_password = os.Getenv("DB_PASSWORD")
		db_username = os.Getenv("DB_USERNAME")
		db_port     = os.Getenv("DB_PORT")
		db_host     = os.Getenv("DB_HOST")
	)

	if db_port == "" {
		db_port = "5432"
	}

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", db_username, db_password, db_host, db_port, db_database)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()

	conn, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		logger.Error("Failed connection to Postgres DB", slog.Any("err", err))
	}

	db := postgres.New(conn)

	tokenAuth := tokenauth.NewTokenAuth(tokenauth.NewTokenAuthParams{
		SecretKey: []byte("secret"),
	})

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			m.TextHTMLMiddleware,
			m.CSPMiddleware,
			jwtauth.Verify(tokenAuth.JWTAuth, TokenFromCookie),
		)

		r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)

		r.Get("/", handlers.NewHomeHandler().ServeHTTP)

		r.Get("/about", handlers.NewAboutHandler().ServeHTTP)

		r.Get("/register", handlers.NewGetRegisterHandler().ServeHTTP)

		r.Post("/register", handlers.NewPostRegisterHandler(handlers.PostRegisterHandlerParams{
			DB: db,
		}).ServeHTTP)

		r.Get("/login", handlers.NewGetLoginHandler().ServeHTTP)

		r.Post("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
			DB:        db,
			TokenAuth: tokenAuth,
		}).ServeHTTP)
	})

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	port := ":8080"

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", port))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")
}
