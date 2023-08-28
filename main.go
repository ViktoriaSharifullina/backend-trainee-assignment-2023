package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testAvito/controllers"
	"testAvito/models"
)

// @title Internship API
// @version 1.0
// @description Swagger API for Internship

func main() {
	route := gin.Default()

	models.InitDB() // new

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	// Создание пользователя
	route.POST("/users", controllers.CreateUser)
	// Получение всех пользователей
	route.GET("/users", controllers.GetUsers)
	// Удаление пользователя
	route.DELETE("/users/:id", controllers.DeleteUser)

	// Создание сегмента
	route.POST("/segments", controllers.CreateSegment)
	// Получение всех сегментов
	route.GET("/segments", controllers.GetSegments)
	// Удаление сегмента
	route.DELETE("/segments/:slug", controllers.DeleteSegment)

	// Добавление пользователя в сегмент
	//route.POST("/segments/:segment_slug/users/:user_id", controllers.AddUserToSegment)

	// Получение активных сегментов пользователя
	route.GET("/users/:user_id/segments", controllers.GetUserSegments)
	// Обновление сегментов пользователя
	route.PUT("/users/:user_id/segments", controllers.UpdateUserSegments)

	// Получение отчета об изменении сегментов упользователя
	route.GET("/history-report", controllers.GenerateSegmentHistoryReport)

	err := route.Run(":8080")
	if err != nil {
		return
	}
}
