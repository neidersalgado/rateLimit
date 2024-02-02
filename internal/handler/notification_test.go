// internal/handler/notification_test.go
package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	custErr "github.com/neidersalgado/rateLimit/internal/error"
	"github.com/neidersalgado/rateLimit/internal/handler"
	"github.com/neidersalgado/rateLimit/internal/service/mocks"
)

func TestNotificationHandler_HandleNotification(t *testing.T) {
	testCases := []struct {
		name                 string
		requestBody          map[string]string
		mockSendNotification func(notificationType, userID, message string) error
		expectedStatusCode   int
		expectedResponse     map[string]string
	}{
		{
			name: "Success",
			requestBody: map[string]string{
				"type":    "status",
				"userId":  "user123",
				"message": "Hello World!",
			},
			mockSendNotification: func(notificationType, userID, message string) error {
				return nil
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: map[string]string{
				"status": "Notification sent",
			},
		},
		{
			name: "RateLimitError",
			requestBody: map[string]string{
				"type":    "status",
				"userId":  "user123",
				"message": "Hello World!",
			},
			mockSendNotification: func(notificationType, userID, message string) error {
				return custErr.NewRateLimitError("Rate limit exceeded", fmt.Errorf("RateLimit Error"))
			},
			expectedStatusCode: http.StatusTooManyRequests,
			expectedResponse: map[string]string{
				"error": "Rate limit exceeded",
			},
		},
		{
			name:        "ValidationError",
			requestBody: map[string]string{},
			mockSendNotification: func(notificationType, userID, message string) error {
				return custErr.NewValidationError("Data Not Valid", fmt.Errorf("validation Error"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: map[string]string{
				"error": "Invalid request",
			},
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.New()

			mockService := mocks.NotificationServiceMock{
				SendNotificationFunc: tc.mockSendNotification,
			}

			notificationHandler := handler.NewNotificationHandler(&mockService)
			router.POST("/notification", notificationHandler.HandleNotification)

			bodyBytes, _ := json.Marshal(tc.requestBody)

			req, _ := http.NewRequest(http.MethodPost, "/notification", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatusCode, w.Code)
			}

			var response map[string]string
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatal("Failed to parse response body")
			}

			if msg, ok := response["error"]; ok && msg != tc.expectedResponse["error"] {
				t.Errorf("Expected error message '%s', got '%s'", tc.expectedResponse["error"], msg)
			}

			if status, ok := response["status"]; ok && status != tc.expectedResponse["status"] {
				t.Errorf("Expected status message '%s', got '%s'", tc.expectedResponse["status"], status)
			}
		})
	}
}
