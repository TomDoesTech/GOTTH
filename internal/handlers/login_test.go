package handlers

import (
	"bytes"

	hashmock "goth/internal/hash/mock"
	"goth/internal/store"
	storemock "goth/internal/store/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {

	testCases := []struct {
		name                string
		email               string
		password            string
		expectedStatusCode  int
		getUserResult       *store.User
		createSessionResult *store.Session
	}{
		{
			name:                "success",
			email:               "test@example.com",
			password:            "password",
			getUserResult:       &store.User{ID: 1, Email: "test@example.com", Password: "password"},
			createSessionResult: &store.Session{UserID: 1, SessionID: "sessionId"},
			expectedStatusCode:  http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			userStore := &storemock.UserStoreMock{}
			sessionStore := &storemock.SessionStoreMock{}
			passwordHash := &hashmock.PasswordHashMock{}

			userStore.On("GetUser", tc.email).Return(tc.getUserResult, nil)

			passwordHash.On("ComparePasswordAndHash", tc.password, tc.getUserResult.Password).Return(true, nil)

			sessionStore.On("CreateSession", &store.Session{UserID: tc.getUserResult.ID}).Return(tc.createSessionResult, nil)

			handler := NewPostLoginHandler(PostLoginHandlerParams{
				UserStore:         userStore,
				SessionStore:      sessionStore,
				PasswordHash:      passwordHash,
				SessionCookieName: "session",
			})
			body := bytes.NewBufferString("email=" + tc.email + "&password=" + tc.password)
			req, _ := http.NewRequest("POST", "/", body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			assert.Equal(tc.expectedStatusCode, rr.Code)

		})
	}
}
