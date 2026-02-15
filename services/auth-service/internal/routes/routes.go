package routes

import (
	"auth-service/internal/controllers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"
	consumer "auth-service/internal/events/consumer"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, rabbitCh *amqp.Channel) {
	api := router.Group("/")

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService, rabbitCh)

	consumer.ConsumeUserUpdated(rabbitCh, authService)

	api.GET("/health", authController.HealthCheck)
	api.POST("/register", authController.Register)
	api.POST("/login", authController.Login)
	api.GET("/refresh", authController.Refresh)
	api.POST("/logout", authController.Logout)
}
