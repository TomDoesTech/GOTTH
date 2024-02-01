package handlers

import (
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
)

type PostLoginHander struct {
	userStore store.UserStore
}

type PostLoginHanderParams struct {
	UserStore store.UserStore
}

func NewPostLoginHandler(params PostLoginHanderParams) *PostLoginHander {
	return &PostLoginHander{
		userStore: params.UserStore,
	}
}

func (h *PostLoginHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.userStore.GetUser(email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	if user.Password == password {
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	c := templates.LoginError()
	c.Render(r.Context(), w)
}
