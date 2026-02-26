package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"user-management-api/structs"
)

// In-memory storage
var Users = []structs.User{}
var IDCounter = 1

// CREATE USER
func CreateUser(c *gin.Context) {

	var payload structs.CreateUserPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	newUser := structs.User{
		ID:   IDCounter,
		Name: payload.Name,
		Age:  payload.Age,
	}

	IDCounter++
	Users = append(Users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

// GET ALL USERS
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
}

// GET USER BY ID
func GetUserByID(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	for _, user := range Users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

// UPDATE USER
func UpdateUser(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var payload structs.UpdateUserPayload

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	for i, user := range Users {
		if user.ID == id {

			Users[i].Name = payload.Name
			Users[i].Age = payload.Age

			c.JSON(http.StatusOK, Users[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

// DELETE USER
func DeleteUser(c *gin.Context) {

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	for i, user := range Users {
		if user.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}