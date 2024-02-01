package strategies

import (
	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"sync"
	"time"
)

type NewsNotificationStrategy struct {
	limit      int
	period     time.Duration
	mutex      sync.RWMutex
	repository Repository
	logg       *logger.ZapLogger
}

func NewNewsNotificationStrategy(c *gin.Context, limit int, period time.Duration, repo Repository) *NewsNotificationStrategy {
	logger, _ := logger.GetLoggerFromContext(c)
	return &NewsNotificationStrategy{
		limit:      limit,
		period:     period,
		repository: repo,
		logg:       logger,
	}
}

func (s *NewsNotificationStrategy) CheckLimitAndSend(userID string) bool {
	return true
}
