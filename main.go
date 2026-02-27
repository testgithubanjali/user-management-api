package main

import (
	"github.com/gin-gonic/gin"
	"user-management-api/database"
	"user-management-api/routes"
)

func main() {

	// Connect to MongoDB
	database.ConnectDB()

	// Create router
	router := gin.Default()

	// Register routes
	routes.UserRoutes(router)

	// Start server
	router.Run(":8080")
}