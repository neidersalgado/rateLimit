package handler

import (
	"github.com/gin-gonic/gin"
	customErr "github.com/neidersalgado/rateLimit/internal/error"
	"github.com/neidersalgado/rateLimit/internal/request"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"net/http"
)

type Service interface {
	SendNotification(*gin.Context, string, string, string) error
}

type NotificationHandler struct {
	srv Service
}

func NewNotificationHandler(srv Service) *NotificationHandler {
	return &NotificationHandler{srv: srv}
}

// HandleNotification maneja la solicitud POST para el endpoint de notificaciones
func (h *NotificationHandler) HandleNotification(c *gin.Context) {
	logger, ok := logger.GetLoggerFromContext(c)
	if !ok {
		return // El error ya se manejó dentro de GetLoggerFromContext
	}
	var req request.NotificationRequest
	logger.Info("Processing notification request")
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.srv.SendNotification(c, req.Type, req.UserID, req.Message); err != nil {
		if customErr, ok := err.(*customErr.CustomError); ok {
			handleCustomError(c, customErr)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Notification sent"})
}

// handleCustomError maneja los errores personalizados y configura la respuesta HTTP adecuada
func handleCustomError(c *gin.Context, err *customErr.CustomError) {
	switch err.Type {
	case customErr.ValidationError:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
	case customErr.RateLimitError:
		c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Message})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
	}
}
