package handlers

import (
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
)

type PostRegisterHander struct {
	userStore store.UserStore
}

type PostRegisterHanderParams struct {
	UserStore store.UserStore
}

func NewPostRegisterHandler(params PostRegisterHanderParams) *PostRegisterHander {
	return &PostRegisterHander{
		userStore: params.UserStore,
	}
}

func (h *PostRegisterHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := h.userStore.CreateUser(email, password)

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	c := templates.RegisterSucces()
	err = c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
