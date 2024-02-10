package handlers

import (
	"goth/internal/auth"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"
	"time"
)

type PostLoginHandler struct {
	userStore store.UserStore
	tokenAuth auth.TokenAuth
}

type PostLoginHandlerParams struct {
	UserStore store.UserStore
	TokenAuth auth.TokenAuth
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore: params.UserStore,
		tokenAuth: params.TokenAuth,
	}
}

func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

		token, err := h.tokenAuth.GenerateToken(*user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "access_token", Value: token, Expires: expiration, Path: "/"}

		http.SetCookie(w, &cookie)

		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	c := templates.LoginError()
	c.Render(r.Context(), w)
}
