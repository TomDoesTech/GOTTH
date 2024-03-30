package mock

import "github.com/stretchr/testify/mock"

type PasswordHashMock struct {
	mock.Mock
}

func (m *PasswordHashMock) ComparePasswordAndHash(password string, encodedHash string) (bool, error) {
	args := m.Called(password, encodedHash)
	return args.Bool(0), args.Error(1)
}

func (m *PasswordHashMock) GenerateFromPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}
