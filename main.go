package main

import (
	"github.com/gin-gonic/gin"
	"testAvito/controllers"
	"testAvito/models"
)

func main() {
	route := gin.Default()

	models.InitDB()

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

	// Получение активных сегментов пользователя
	route.GET("/users/:user_id/segments", controllers.GetUserSegments)
	// Обновление сегментов пользователя
	route.PUT("/users/:user_id/segments", controllers.UpdateUserSegments)

	// Получение отчета об изменении сегментов упользователя
	route.GET("/history-report", controllers.GenerateSegmentHistoryReport)

	db := models.GetDB() // Получение подключения к базе данных

	go controllers.StartExpirationChecker(db)

	err := route.Run(":8080")
	if err != nil {
		return
	}
}
