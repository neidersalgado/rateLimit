package mocks

import (
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetUserLimits(userID string) (map[string]repository.Limit, bool) {
	args := m.Called(userID)
	return args.Get(0).(map[string]repository.Limit), args.Bool(1)
}

func (m *RepositoryMock) UpdateUserLimit(userID string, notificationType string, newLimit repository.Limit) {
	m.Called(userID, notificationType, newLimit)
}
