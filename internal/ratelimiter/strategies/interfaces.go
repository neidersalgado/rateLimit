package strategies

import "github.com/neidersalgado/rateLimit/internal/repository"

type Repository interface {
	GetUserLimits(userID string) (map[string]repository.Limit, bool)
	UpdateUserLimit(userID string, notificationType string, newLimit repository.Limit)
}
