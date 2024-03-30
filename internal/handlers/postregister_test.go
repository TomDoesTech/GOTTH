package handlers

import (
	"bytes"

	storemock "goth/internal/store/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRegisterUserHandler(t *testing.T) {

	testCases := []struct {
		name               string
		email              string
		password           string
		createUserError    error
		expectedStatusCode int
		expectedBody       []byte
	}{
		{
			name:               "success",
			email:              "test@example.com",
			password:           "password",
			expectedStatusCode: http.StatusOK,
			expectedBody:       []byte(`<h1>Registration successful</h1><p>Go to <a href="login">login</a></p>`),
		},
		{
			name:               "fail - error creating user",
			email:              "test@example.com",
			password:           "password",
			createUserError:    gorm.ErrDuplicatedKey,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       []byte(`<h1>Registration failed</h1>`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			userStore := &storemock.UserStoreMock{}

			userStore.On("CreateUser", tc.email, tc.password).Return(tc.createUserError)

			handler := NewPostRegisterHandler(PostRegisterHandlerParams{
				UserStore: userStore,
			})
			body := bytes.NewBufferString("email=" + tc.email + "&password=" + tc.password)
			req, _ := http.NewRequest("POST", "/", body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			assert.Equal(tc.expectedStatusCode, rr.Code, "handler returned wrong status code: got %v want %v", rr.Code, tc.expectedStatusCode)

			assert.True(bytes.Contains(rr.Body.Bytes(), tc.expectedBody), "handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)

			userStore.AssertExpectations(t)
		})
	}
}
