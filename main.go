package main

import (
	"net/http"
	"github.com/SharpDenin/PrBack/config"
	"github.com/SharpDenin/PrBack/controllers"
	// "github.com/SharpDenin/PrBack/models"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	config.ConnectToDB()

	route.GET("/users", controllers.GetAllUsers)
	route.POST("/users", controllers.CreateUser)
	route.GET("/users/:id", controllers.GetUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run()
}