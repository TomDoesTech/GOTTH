package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"goth/templates"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/"+tmplName,
		"templates/partial/header.html",
		"templates/partial/footer.html",
		"templates/partial/base.html",
	)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

type User struct {
	Email    string
	Password string
}

func generateRandomString(length int) string {

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func cspMiddleware(next http.Handler) http.Handler {
	htmxNonce := generateRandomString(16)
	twNonce := generateRandomString(16)

	// set then in context
	ctx := context.WithValue(context.Background(), "htmxNonce", htmxNonce)
	ctx = context.WithValue(ctx, "twNonce", twNonce)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'"+" script-src 'nonce-"+htmxNonce+"'"+""+"style-src"+" 'nonce-"+twNonce+"'")

		next.ServeHTTP(w, r)
	})

}

func main() {

	users := []User{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html charset=utf-8")
				next.ServeHTTP(w, r)
			})
		})

		r.Use(cspMiddleware)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			c := templates.Index()
			err := templates.Layout(c, "My website").Render(context.Background(), w)

			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		r.Get("/about", func(w http.ResponseWriter, r *http.Request) {

			component := templates.About("About")
			err := component.Render(context.Background(), w)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		r.Get("/register", func(w http.ResponseWriter, r *http.Request) {

			component := templates.Register("Resgiter")
			err := component.Render(context.Background(), w)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			email := r.FormValue("email")
			password := r.FormValue("password")

			users = append(users, User{Email: email, Password: password})
			w.WriteHeader(http.StatusCreated)

			fmt.Fprintf(w, "<h1>Registration successful</h1><p>Go to <a href=\"/login\">login</a></p>")
		})

		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			component := templates.Login("Login")
			err := component.Render(context.Background(), w)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			email := r.FormValue("email")
			password := r.FormValue("password")

			for _, user := range users {
				if user.Email == email && user.Password == password {
					w.WriteHeader(http.StatusOK)

					fmt.Fprintf(w, "<h1>Login successful</h1><p>Go to <a href=\"/\">home</a></p>")
					return
				}
			}

			w.WriteHeader(http.StatusUnauthorized)

			fmt.Fprintf(w, "<h1>Unauthorized</h1><p>Go to <a href=\"/login\">login</a></p>")
		})

	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		RenderTemplate(w, "404.html", nil, r)
	})

	err := http.ListenAndServe(":8080", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
