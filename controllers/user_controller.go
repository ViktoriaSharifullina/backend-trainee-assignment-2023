package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testAvito/models"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := models.GetDB() // Получение подключения к базе данных
	user := models.User{Username: input.Username}
	if err := models.CreateUser(db, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user_db := db.First(&user, user.ID)

	c.JSON(http.StatusOK, gin.H{"data": user_db.Value})
}

func GetUsers(c *gin.Context) {
	db := models.GetDB()

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	db := models.GetDB()
	userID := c.Param("id") // Получаем идентификатор пользователя из URL

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
