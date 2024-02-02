package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck maneja la solicitud GET para el endpoint de health check
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Service is up and running",
	})
}
