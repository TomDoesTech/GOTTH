package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type GetLoginHander struct{}

func NewGetLoginHandler() *GetLoginHander {
	return &GetLoginHander{}
}

func (h *GetLoginHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Login("Login")
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
