package repository

import (
	"sync"
	"time"
)

type Limit struct {
	Count            int
	LastNotification time.Time
}

// UserRepo gestiona los límites de notificaciones para cada usuario y tipo.
type UserRepo struct {
	UserLimits map[string]map[string]Limit
	mutex      sync.RWMutex
}

// NewUserRepo crea una nueva instancia de UserRepo.
func NewUserRepo() *UserRepo {
	return &UserRepo{
		UserLimits: make(map[string]map[string]Limit),
	}
}

// GetUserLimits retorna los límites para un usuario específico.
func (r *UserRepo) GetUserLimits(userID string) (map[string]Limit, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	limits, exists := r.UserLimits[userID]
	return limits, exists
}

// UpdateUserLimit actualiza el límite para un usuario y tipo específicos.
func (r *UserRepo) UpdateUserLimit(userID string, notificationType string, newLimit Limit) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.UserLimits[userID]; !exists {
		r.UserLimits[userID] = make(map[string]Limit)
	}
	r.UserLimits[userID][notificationType] = newLimit
}
