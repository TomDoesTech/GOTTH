package router

import (
	"goth/internal/config"
	"goth/internal/handlers"
	"goth/internal/hash/passwordhash"
	"goth/internal/middleware"
	"goth/internal/store/dbstore"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

// Define a struct to group dependencies
type RouterDependencies struct {
	Config         config.Config
	UserStore      *dbstore.UserStore
	SessionStore   *dbstore.SessionStore
	PasswordHasher *passwordhash.PasswordHash
}

func SetupRouter(deps RouterDependencies) *chi.Mux {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	authMiddleware := middleware.NewAuthMiddleware(deps.SessionStore, deps.Config.SessionCookieName)

	r.Group(func(r chi.Router) {
		r.Use(
			chiMiddleware.Logger,
			middleware.TextHTMLMiddleware,
			middleware.CSPMiddleware,
			authMiddleware.AddUserToContext,
		)

		r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)

		r.Get("/", handlers.NewHomeHandler().ServeHTTP)

		r.Get("/about", handlers.NewAboutHandler().ServeHTTP)

		r.Get("/register", handlers.NewGetRegisterHandler().ServeHTTP)

		r.Post("/register", handlers.NewPostRegisterHandler(handlers.PostRegisterHandlerParams{
			UserStore: deps.UserStore,
		}).ServeHTTP)

		r.Get("/login", handlers.NewGetLoginHandler().ServeHTTP)

		r.Post("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
			UserStore:         deps.UserStore,
			SessionStore:      deps.SessionStore,
			PasswordHash:      deps.PasswordHasher,
			SessionCookieName: deps.Config.SessionCookieName,
		}).ServeHTTP)

		r.Post("/logout", handlers.NewPostLogoutHandler(handlers.PostLogoutHandlerParams{
			SessionCookieName: deps.Config.SessionCookieName,
		}).ServeHTTP)
	})

	return r
}
