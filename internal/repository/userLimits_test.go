package repository_test

import (
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewUserRepo(t *testing.T) {
	repo := repository.NewUserRepo()
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.UserLimits)
}

func TestUserRepo_GetUserLimits(t *testing.T) {
	repo := repository.NewUserRepo()
	userID := "user1"
	notificationType := "Status"
	limit := repository.Limit{Count: 2, LastNotification: time.Now()}

	repo.UserLimits[userID] = map[string]repository.Limit{notificationType: limit}

	tests := []struct {
		name           string
		userID         string
		expectedLimits map[string]repository.Limit
		expectedExists bool
	}{
		{
			name:           "existing user",
			userID:         userID,
			expectedLimits: map[string]repository.Limit{notificationType: limit},
			expectedExists: true,
		},
		{
			name:           "non-existing user",
			userID:         "user2",
			expectedLimits: nil,
			expectedExists: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			limits, exists := repo.GetUserLimits(tc.userID)

			assert.Equal(t, tc.expectedExists, exists, "exists should match expected value")
			if exists {
				assert.Equal(t, tc.expectedLimits, limits, "limits should match expected value")
			}
		})
	}
}

func TestUserRepo_UpdateUserLimit(t *testing.T) {
	repo := repository.NewUserRepo()
	userID := "user1"
	notificationType := "Status"
	newLimit := repository.Limit{Count: 5, LastNotification: time.Now()}

	repo.UpdateUserLimit(userID, notificationType, newLimit)

	limits, exists := repo.GetUserLimits(userID)
	assert.True(t, exists, "user should exist after update")
	assert.NotNil(t, limits[notificationType], "limit should be set for notification type")
	assert.Equal(t, newLimit, limits[notificationType], "limit should match new limit")
}
