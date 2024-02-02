package strategies

import (
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/pkg/timeUtil"
	"sync"
	"time"
)

type MarketingNotificationStrategy struct {
	limit        int
	period       time.Duration
	mutex        sync.RWMutex
	repository   Repository
	timeProvider timeUtil.TimeProvider
}

func NewMarketingNotificationStrategy(limit int, period time.Duration, repo Repository, timeProv timeUtil.TimeProvider) *MarketingNotificationStrategy {
	return &MarketingNotificationStrategy{
		limit:        limit,
		period:       period,
		repository:   repo,
		timeProvider: timeProv,
	}
}

func (s *MarketingNotificationStrategy) CheckLimitAndSend(userID string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userLimits, exists := s.repository.GetUserLimits(userID)
	if !exists {
		s.repository.UpdateUserLimit(userID, "Marketing", repository.Limit{Count: 1, LastNotification: s.timeProvider.Now()})
		return true
	}

	limit, exists := userLimits["Marketing"]
	if !exists {
		s.repository.UpdateUserLimit(userID, "Marketing", repository.Limit{Count: 1, LastNotification: s.timeProvider.Now()})
		return true
	}

	if time.Since(limit.LastNotification) > s.period {
		s.repository.UpdateUserLimit(userID, "Marketing", repository.Limit{Count: 1, LastNotification: s.timeProvider.Now()})
		return true
	} else if limit.Count < s.limit {
		s.repository.UpdateUserLimit(userID, "Marketing", repository.Limit{Count: limit.Count + 1, LastNotification: limit.LastNotification})
		return true
	}

	return false
}
