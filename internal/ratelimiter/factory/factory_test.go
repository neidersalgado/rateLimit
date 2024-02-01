package factory_test

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/internal/ratelimiter/factory"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/internal/repository/mocks"
	"github.com/neidersalgado/rateLimit/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http/httptest"
)

func TestCreateNotificationStrategy(t *testing.T) {
	// Crear e inicializar el logger
	zapLogger, _ := zap.NewDevelopment()
	testLogger := &logger.ZapLogger{Logger: zapLogger}

	// Crear el contexto de Gin y añadir el logger
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Set("logger", testLogger)

	// Crear y configurar el mock de Repository
	mockRepo := &mocks.RepositoryMock{}
	mockRepo.On("GetUserLimits", mock.Anything).Return(make(map[string]repository.Limit), true)
	mockRepo.On("UpdateUserLimit", mock.Anything, mock.Anything, mock.Anything).Return()

	// Definir los casos de prueba
	testCases := []struct {
		name             string
		notificationType string
		wantErr          bool
	}{
		// ... tus casos de prueba
	}

	// Ejecutar los casos de prueba
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			strategy, err := factory.CreateNotificationStrategy(ctx, tc.notificationType, mockRepo)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, strategy)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, strategy)
			}
		})
	}
}
