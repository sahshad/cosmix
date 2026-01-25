package routes

import (
	"user-service/internal/controllers"
	"user-service/internal/events"
	"user-service/internal/repositories"
	"user-service/internal/services"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, rabbitCh *amqp.Channel) {
	api := router.Group("/")

	profileRepo := repositories.NewUserProfileRepository(db)
	profileService := services.NewUserProfileService(profileRepo)
	profileController := controllers.NewUserProfileController(profileService)

	followRepo := repositories.NewFollowRepository(db)
	followService := services.NewFollowService(followRepo)
	followController := controllers.NewFollowController(followService)
	
	events.ConsumeUserCreated(rabbitCh, profileService)

	api.GET("/health", profileController.HealthCheck)
	api.GET("/me", profileController.GetMe)
	api.PUT("/me", profileController.UpdateMe)
	api.GET("/username/:username", profileController.GetByUsername)
	api.POST("/:id/follow", followController.Follow)
	api.DELETE("/:id/unfollow", followController.Unfollow)
	api.GET("/:id/followers", followController.GetFollowers)
	api.GET("/:id/following", followController.GetFollowing)
}
