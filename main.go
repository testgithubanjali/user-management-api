package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"user-management-api/database"
	"user-management-api/routes"
)

func main() {

	// Load .env file (only for development)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to MongoDB
	database.ConnectDB()

	// Create router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.UserRoutes(router)

	// Get PORT from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	log.Println("Server running on port:", port)

	// Start server
	router.Run(":" + port)
}