package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neidersalgado/rateLimit/internal/handler"
	"github.com/neidersalgado/rateLimit/internal/repository"
	"github.com/neidersalgado/rateLimit/internal/service"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	repo := repository.NewUserRepo()
	NotificationSrv := service.NewNotificationService(repo)
	notificationHandler := handler.NewNotificationHandler(NotificationSrv)
	router.GET("/healt", handler.HealthCheck)
	router.POST("/notification", notificationHandler.HandleNotification)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Starting server on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("ListenAndServe:", zap.Error(err))
	}
}
