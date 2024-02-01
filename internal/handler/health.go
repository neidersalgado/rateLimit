package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"net/http"
)

// HealthCheck maneja la solicitud GET para el endpoint de health check
func HealthCheck(c *gin.Context) {
	zaplog, _ := logger.GetLoggerFromContext(c)
	zaplog.Info("Processing notification healt")
	c.JSON(http.StatusOK, gin.H{
		"status": "Service is up and running",
	})
}
