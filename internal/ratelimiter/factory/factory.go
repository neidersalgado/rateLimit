package factory

import (
	"fmt"
	"github.com/neidersalgado/rateLimit/internal/ratelimiter/strategies"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/pkg/timeUtil"
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

func CreateNotificationStrategy(notificationType string, repo Repository) (NotificationStrategy, error) {
	timeProvider := &timeUtil.DefaultTimeProvider{}
	switch notificationType {
	case "Status":
		return strategies.NewStatusNotificationStrategy(2, 1*time.Minute, repo, timeProvider), nil
	case "News":
		return strategies.NewNewsNotificationStrategy(2, 1*time.Hour, repo, timeProvider), nil
	case "Marketing":
		return strategies.NewMarketingNotificationStrategy(2, time.Hour*24*30, repo, timeProvider), nil
	default:
		return nil, fmt.Errorf("not Valid Notification Type")
	}
}
