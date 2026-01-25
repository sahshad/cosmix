package main

import (
	"log"
	"os"

	"user-service/internal/database"
	"user-service/internal/events"
	"user-service/internal/routes"

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
		port = "8081"
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
	rabbitChannel := events.NewRabbitMQChannel(rabbitURL)
	if rabbitChannel == nil {
		log.Println("RabbitMQ unavailable, running without consumer")
	}

	// HTTP router
	router := gin.Default()
	routes.RegisterRoutes(router, db, rabbitChannel)

	log.Println("User service running on :" + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
