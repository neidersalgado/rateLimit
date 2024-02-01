package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	custErr "github.com/neidersalgado/rateLimit/internal/error"
	"github.com/neidersalgado/rateLimit/internal/ratelimiter/factory"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/pkg/logger"
)

type Strategy interface {
}

type Repository interface {
	GetUserLimits(userID string) (map[string]repository.Limit, bool)
	UpdateUserLimit(userID string, notificationType string, newLimit repository.Limit)
}

type NotificationService struct {
	Strategy Strategy
	Repo     Repository
}

func NewNotificationService(repo Repository) *NotificationService {
	return &NotificationService{Repo: repo}
}

func (s *NotificationService) SendNotification(ctx *gin.Context, notificationType, userID, message string) error {
	logger, _ := logger.GetLoggerFromContext(ctx)
	logger.Debug("Send Notification Service")
	NotificationStrategy, err := factory.CreateNotificationStrategy(ctx, notificationType, s.Repo)
	if err != nil {
		return custErr.NewInvalidRequest("invalid Notification Type", err)
	}
	if check := NotificationStrategy.CheckLimitAndSend(userID); !check {
		return custErr.NewRateLimitError("Rate Validation Error", fmt.Errorf("Max notification send, ty later"))
	}
	return nil
}
