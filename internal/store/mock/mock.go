package mock

import (
	"goth/internal/store"

	"github.com/stretchr/testify/mock"
)

type UserStoreMock struct {
	mock.Mock
}

func (m *UserStoreMock) CreateUser(email string, password string) error {
	args := m.Called(email, password)

	return args.Error(0)
}

func (m *UserStoreMock) GetUser(email string) (*store.User, error) {
	args := m.Called(email)
	return args.Get(0).(*store.User), args.Error(1)
}

type SessionStoreMock struct {
	mock.Mock
}

func (m *SessionStoreMock) CreateSession(session *store.Session) (*store.Session, error) {
	args := m.Called(session)
	return args.Get(0).(*store.Session), args.Error(1)
}

func (m *SessionStoreMock) GetUserFromSession(sessionID string, userID string) (*store.User, error) {
	args := m.Called(sessionID, userID)
	return args.Get(0).(*store.User), args.Error(1)
}
