package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testAvito/models"
)

func main() {
	route := gin.Default()

	models.ConnectDB() // new

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run()
}
