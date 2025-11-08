package routes

import (
	"auth-service/internal/controllers"
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api/auth")

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	api.POST("/register", authController.Register)
	api.POST("/login", authController.Login)
	api.GET("/refresh", authController.Refresh)
	api.POST("/logout", authController.Logout)
}
