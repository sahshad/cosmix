package routes

import (
	app "auth-service/internal/app"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, container *app.Container) {
	api := router.Group("/")

	api.GET("/health", container.AuthController.HealthCheck)
	api.POST("/register", container.AuthController.Register)
	api.POST("/login", container.AuthController.Login)
	api.GET("/refresh", container.AuthController.Refresh)
	api.POST("/logout", container.AuthController.Logout)
}
