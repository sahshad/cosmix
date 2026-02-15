package routes

import (
	"user-service/internal/app"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, container *app.Container) {
	api := router.Group("/")

	api.GET("/health", container.UserProfileController.HealthCheck)
	api.GET("/me", container.UserProfileController.GetMe)
	api.PUT("/me", container.UserProfileController.UpdateMe)
	api.GET("/username/:username", container.UserProfileController.GetByUsername)
	api.POST("/:id/follow", container.FollowController.Follow)
	api.DELETE("/:id/unfollow", container.FollowController.Unfollow)
	api.GET("/:id/followers", container.FollowController.GetFollowers)
	api.GET("/:id/following", container.FollowController.GetFollowing)
}
