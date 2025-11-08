// package main

// import (
// 	"log"

// 	"auth-service/internal/database"
// 	"auth-service/internal/routes"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("⚠️  No .env file found, using environment variables")
// 	}

// 	db, err := database.ConnectDB()
// 	if err != nil {
// 		log.Fatalf("Database connection failed: %v", err)
// 	}

// 	r := gin.Default()
// 	routes.SetupRoutes(r, db)

// 	log.Println("🚀 Auth service running on :8080")
// 	if err := r.Run(":8080"); err != nil {
// 		log.Fatalf("Server failed to start: %v", err)
// 	}
// }

package main

import (
	"log"
	"os"

	"auth-service/internal/database"
	"auth-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// optionally override Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}

	r := gin.Default()
	// (CORS or other middleware can be added here)

	routes.RegisterRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Auth service running on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
