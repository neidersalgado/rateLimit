package strategies_test

import (
	"github.com/neidersalgado/rateLimit/internal/ratelimiter/strategies"
	"github.com/neidersalgado/rateLimit/internal/repository"
	repoMock "github.com/neidersalgado/rateLimit/internal/repository/mocks"
	timeMock "github.com/neidersalgado/rateLimit/pkg/timeUtil/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestStatusNotificationStrategy_CheckLimitAndSend(t *testing.T) {
	mockTime := time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)
	justBeforeLimit := mockTime.Add(-59 * time.Second)
	justAfterLimit := mockTime.Add(1 * time.Minute)

	testCases := []struct {
		name              string
		existingLimits    map[string]repository.Limit
		expectedResult    bool
		expectedUpdate    repository.Limit
		timeNow           time.Time
		callUpdate        bool
		callGetUserLimits bool
	}{
		{
			name: "Within limit and period",
			existingLimits: map[string]repository.Limit{
				"Status": {
					Count:            1,
					LastNotification: justBeforeLimit,
				},
			},
			expectedResult:    true,
			expectedUpdate:    repository.Limit{Count: 2, LastNotification: justBeforeLimit},
			timeNow:           mockTime,
			callUpdate:        true,
			callGetUserLimits: true,
		},
		{
			name:              "New limit - no existing user limits",
			existingLimits:    nil,
			expectedResult:    true,
			expectedUpdate:    repository.Limit{Count: 1, LastNotification: mockTime},
			timeNow:           mockTime,
			callUpdate:        true,
			callGetUserLimits: true,
		},
		{
			name: "Limit refreshes after period",
			existingLimits: map[string]repository.Limit{
				"Status": {
					Count:            2,
					LastNotification: justBeforeLimit,
				},
			},
			expectedResult:    true,
			expectedUpdate:    repository.Limit{Count: 1, LastNotification: justAfterLimit},
			timeNow:           justAfterLimit,
			callUpdate:        true,
			callGetUserLimits: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(repoMock.RepositoryMock)
			mockTimeProvider := new(timeMock.MockTimeProvider)
			mockTimeProvider.On("Now").Return(tc.timeNow)

			if tc.callGetUserLimits {
				mockRepo.On("GetUserLimits", mock.Anything).Return(tc.existingLimits, true)
			}
			if tc.callUpdate {
				mockRepo.On("UpdateUserLimit", mock.Anything, "Status", mock.AnythingOfType("repository.Limit")).Return()
			}

			strategy := strategies.NewStatusNotificationStrategy(2, 1*time.Minute, mockRepo, mockTimeProvider)
			result := strategy.CheckLimitAndSend("userID")

			assert.Equal(t, tc.expectedResult, result)
			mockRepo.AssertExpectations(t)
			mockTimeProvider.AssertExpectations(t)
		})
	}

}
