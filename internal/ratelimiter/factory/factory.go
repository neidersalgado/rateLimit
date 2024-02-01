package factory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/internal/ratelimiter/strategies"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"time"
)

// NotificationStrategy define la interfaz para las estrategias de notificación
type NotificationStrategy interface {
	CheckLimitAndSend(string) bool
}

type Repository interface {
	GetUserLimits(userID string) (map[string]repository.Limit, bool)
	UpdateUserLimit(userID string, notificationType string, newLimit repository.Limit)
}

func CreateNotificationStrategy(ctx *gin.Context, notificationType string, repo Repository) (NotificationStrategy, error) {
	logger, _ := logger.GetLoggerFromContext(ctx)
	switch notificationType {
	case "Status":
		logger.Info("Build Statu Strategy")
		return strategies.NewStatusNotificationStrategy(ctx, 2, 1*time.Minute, repo), nil
	case "News":
		logger.Info("Build News Strategy")
		return strategies.NewNewsNotificationStrategy(ctx, 2, 1*time.Hour, repo), nil
	case "Marketing":
		logger.Info("Build Marketing Strategy")
		return strategies.NewMarketingNotificationStrategy(ctx, 2, time.Hour*24*30, repo), nil
	default:
		return nil, fmt.Errorf("not Valid Notification Type")
	}
}
