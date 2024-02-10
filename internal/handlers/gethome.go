package handlers

import (
	"goth/internal/templates"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type HomeHander struct{}

func NewHomeHandler() *HomeHander {
	return &HomeHander{}
}

func (h *HomeHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, claims, _ := jwtauth.FromContext(r.Context())

	email, ok := claims["email"].(string)

	if !ok {
		c := templates.GuestIndex()
		err := templates.Layout(c, "My website").Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		return
	}

	c := templates.Index(email)
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
