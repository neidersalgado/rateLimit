package mocks

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type MockTimeProvider struct {
	mock.Mock
}

func (m *MockTimeProvider) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
