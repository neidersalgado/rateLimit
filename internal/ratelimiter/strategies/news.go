package strategies

import (
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/pkg/timeUtil"
	"sync"
	"time"
)

type NewsNotificationStrategy struct {
	limit        int
	period       time.Duration
	mutex        sync.RWMutex
	repository   Repository
	timeProvider timeUtil.TimeProvider
}

func NewNewsNotificationStrategy(limit int, period time.Duration, repo Repository, timeProv timeUtil.TimeProvider) *NewsNotificationStrategy {
	return &NewsNotificationStrategy{
		limit:        limit,
		period:       period,
		repository:   repo,
		timeProvider: timeProv,
	}
}

func (s *NewsNotificationStrategy) CheckLimitAndSend(userID string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userLimits, exists := s.repository.GetUserLimits(userID)
	if !exists {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: 1, LastNotification: s.timeProvider.Now()})
		return true
	}

	limit, exists := userLimits["Status"]
	if !exists {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: 1, LastNotification: s.timeProvider.Now()})
		return true
	}

	if time.Since(limit.LastNotification) > s.period {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: 1, LastNotification: s.timeProvider.Now()})
		return true
	} else if limit.Count < s.limit {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: limit.Count + 1, LastNotification: limit.LastNotification})
		return true
	}

	return false
}
