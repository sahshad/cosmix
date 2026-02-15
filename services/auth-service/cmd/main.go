package main

import (
	"log"
	"os"

	"auth-service/internal/app"
	"auth-service/internal/database"
	messaging "auth-service/internal/messaging"
	"auth-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	// Database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	// RabbitMQ
	rabbit := messaging.Connect(rabbitURL)
	if rabbit == nil {
		log.Println("RabbitMQ unavailable, events will not be published")
	}

	container := app.NewContainer(db, rabbit)

	// HTTP router
	router := gin.Default()

	routes.RegisterRoutes(router, container)
	app.RegisterConsumers(container)

	log.Println("Auth service running on :" + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
