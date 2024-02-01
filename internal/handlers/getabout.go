package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type AboutHander struct{}

func NewAboutHandler() *AboutHander {
	return &AboutHander{}
}

func (h *AboutHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.About("About")
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
