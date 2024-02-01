package strategies

import (
	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"sync"
	"time"
)

type MarketingNotificationStrategy struct {
	limit      int
	period     time.Duration
	mutex      sync.RWMutex
	repository Repository
	logg       *logger.ZapLogger
}

func NewMarketingNotificationStrategy(c *gin.Context, limit int, period time.Duration, repo Repository) *MarketingNotificationStrategy {
	logger, _ := logger.GetLoggerFromContext(c)
	return &MarketingNotificationStrategy{
		limit:      limit,
		period:     period,
		repository: repo,
		logg:       logger,
	}
}

func (s *MarketingNotificationStrategy) CheckLimitAndSend(userID string) bool {
	s.logg.Info("Check Limit Marketing")

	return true
}
