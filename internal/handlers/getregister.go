package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type GetRegisterHander struct{}

func NewGetRegisterHandler() *GetRegisterHander {
	return &GetRegisterHander{}
}

func (h *GetRegisterHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.RegisterPage()
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
