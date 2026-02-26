package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)
type user struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}
var users= []user{}
var idCounter = 1
func main(){
	router:= gin.Default()
	router.POST("/users", func(c *gin.Context) {
		var newUser user
		if err:= c.BindJSON(&newUser); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"errror": "invalid input",
			})
			return
		}
		newUser.ID = idCounter
		idCounter++
		users=append(users,newUser)
		c.JSON(http.StatusOK, newUser)
	})
	router.GET("/users", func(c *gin.Context){
		c.JSON(http.StatusOK,users)
	})
	router.GET("/users/:id",func(c *gin.Context){
		idParam:= c.Param("id")
		id,_ := strconv.Atoi(idParam)
		for _, user := range users{
			if user.ID ==id{
				c.JSON(http.StatusOK,user)
				return
		
			}
			
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})

	})
}