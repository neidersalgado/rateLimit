package strategies

import (
	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"go.uber.org/zap"
	"sync"
	"time"
)

type StatusNotificationStrategy struct {
	limit      int
	period     time.Duration
	mutex      sync.RWMutex
	repository Repository
	logg       *logger.ZapLogger
}

func NewStatusNotificationStrategy(ctx *gin.Context, limit int, period time.Duration, repo Repository) *StatusNotificationStrategy {
	logger, _ := logger.GetLoggerFromContext(ctx)
	return &StatusNotificationStrategy{
		limit:      limit,
		period:     period,
		repository: repo,
		logg:       logger,
	}
}

func (s *StatusNotificationStrategy) CheckLimitAndSend(userID string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.logg.Info("Check Limit For User", zap.Any("UserID", userID))
	userLimits, exists := s.repository.GetUserLimits(userID)
	if !exists {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: 1, LastNotification: time.Now()})
		return true
	}

	limit, exists := userLimits["Status"]
	if !exists {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: 1, LastNotification: time.Now()})
		return true
	}

	if time.Since(limit.LastNotification) > s.period {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: 1, LastNotification: time.Now()})
		return true
	} else if limit.Count < s.limit {
		s.repository.UpdateUserLimit(userID, "Status", repository.Limit{Count: limit.Count + 1, LastNotification: limit.LastNotification})
		return true
	}

	return false
}
