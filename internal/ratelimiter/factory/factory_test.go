package factory_test

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/neidersalgado/rateLimit/internal/ratelimiter/factory"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateNotificationStrategy(t *testing.T) {
	mockRepo := &mocks.RepositoryMock{}
	mockRepo.On("GetUserLimits", mock.Anything).Return(make(map[string]repository.Limit), true)
	mockRepo.On("UpdateUserLimit", mock.Anything, mock.Anything, mock.Anything).Return()

	testCases := []struct {
		name             string
		notificationType string
		wantErr          bool
	}{
		//TODO
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			strategy, err := factory.CreateNotificationStrategy(tc.notificationType, mockRepo)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, strategy)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, strategy)
			}
		})
	}
}
