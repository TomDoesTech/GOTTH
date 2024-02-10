package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.NotFound()
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
