package strategies

import (
	"sync"
	"time"
)

// userLimitState mantiene el estado de límite de tasa para cada usuario
type userLimitState struct {
	lastNotification  time.Time // La última vez que se envió una notificación
	notificationCount int       // Cuántas notificaciones se han enviado en el período actual
}

type StatusNotificationStrategy struct {
	limit     int
	period    time.Duration
	userState map[string]*userLimitState
	mutex     sync.RWMutex
}

func NewStatusNotificationStrategy(limit int, period time.Duration) *StatusNotificationStrategy {
	return &StatusNotificationStrategy{
		limit:     limit,
		period:    period,
		userState: make(map[string]*userLimitState),
	}
}

func (s *StatusNotificationStrategy) CheckLimitAndSend(userID string) bool {
	return true
}
