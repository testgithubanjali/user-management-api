package main

import (
	"github.com/gin-gonic/gin"
	"user-management-api/routes"
)

func main() {

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run(":8080")
}