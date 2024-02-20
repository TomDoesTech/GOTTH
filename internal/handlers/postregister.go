package handlers

import (
	"goth/internal/db"
	"goth/internal/templates"
	"net/http"
)

type PostRegisterHandler struct {
	DB *db.Queries
}

type PostRegisterHandlerParams struct {
	DB *db.Queries
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		DB: params.DB,
	}
}

func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := h.DB.CreateUser(r.Context(), db.CreateUserParams{
		Email:    email,
		Password: password,
	})

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	c := templates.RegisterSucces()
	errTemplate := c.Render(r.Context(), w)

	if errTemplate != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}
