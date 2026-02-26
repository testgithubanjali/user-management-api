package routes

import (
	"github.com/gin-gonic/gin"
	"user-management-api/handlers"
)

func UserRoutes(router *gin.Engine) {

	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetAllUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
}