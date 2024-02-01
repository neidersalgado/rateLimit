// pkg/Logger/util.go

package logger

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoggerFromContext(c *gin.Context) (*ZapLogger, bool) {
	zapLogger, exists := c.Get("Logger")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return nil, false
	}

	logger, ok := zapLogger.(*ZapLogger)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger found, but casting failed"})
		return nil, false
	}

	return logger, true
}
