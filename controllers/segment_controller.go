package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"testAvito/models"
	"time"
)

type CreateSegmentInput struct {
	Slug string `json:"slug" binding:"required"`
}

func CreateSegment(c *gin.Context) {
	var input CreateSegmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := models.GetDB()
	segment := models.Segment{Slug: input.Slug}
	if err := models.CreateSegment(db, segment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create segment"})
		return
	}

	segment_db := db.First(&segment, segment.ID)

	c.JSON(http.StatusOK, gin.H{"data": segment_db.Value})
}

func GetSegments(c *gin.Context) {
	db := models.GetDB()

	var segments []models.Segment
	if err := db.Find(&segments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch segments"})
		return
	}

	c.JSON(http.StatusOK, segments)
}

func DeleteSegment(c *gin.Context) {
	db := models.GetDB()
	segmentSlug := c.Param("slug")

	var segment models.Segment
	if err := db.Where("slug = ?", segmentSlug).First(&segment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	if err := db.Delete(&segment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment deleted successfully"})
}

type UserSegmentUpdateInput struct {
	AddSegments    []string `json:"add_segments"`
	RemoveSegments []string `json:"remove_segments"`
}

// UpdateUserSegments Метод добавления и удаления сегментов пользователю
func UpdateUserSegments(c *gin.Context) {
	db := models.GetDB()
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input UserSegmentUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.AddSegments) == 0 && len(input.RemoveSegments) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both add_segments and remove_segments cannot be empty"})
		return
	}

	fmt.Println(input)
	for _, segmentSlug := range input.AddSegments {
		var segment models.Segment
		if err := db.Where("slug = ?", segmentSlug).First(&segment).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Segment not found"})
			return
		}

		userSegment := models.UserSegment{UserID: uint(userID), SegmentID: segment.ID}
		if err := models.CreateUserSegment(db, userSegment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to segment"})
			return
		}

		// Запись истории попадания пользователя в сегмент
		historyRecord := models.UserSegmentHistory{
			UserID:    uint(userID),
			SegmentID: segment.ID,
			Operation: "add",
			Date:      time.Now(),
		}
		if err := models.CreateHistory(db, historyRecord); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log segment history"})
			return
		}

	}

	for _, segmentSlug := range input.RemoveSegments {
		var segment models.Segment
		if err := db.Where("slug = ?", segmentSlug).First(&segment).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Segment not found"})
			return
		}

		if err := db.Where("user_id = ? AND segment_id = ?", user.ID, segment.ID).Delete(models.UserSegment{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove segment from user"})
			return
		}

		// Запись истории выбывания пользователя из сегмента
		historyRecord := models.UserSegmentHistory{
			UserID:    uint(userID),
			SegmentID: segment.ID,
			Operation: "remove",
			Date:      time.Now(),
		}
		if err := models.CreateHistory(db, historyRecord); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log segment history"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User segments updated successfully"})
}

func GetUserSegments(c *gin.Context) {
	db := models.GetDB()
	userID := c.Param("user_id")

	var userSegments []models.UserSegment
	if err := db.Where("user_id = ?", userID).Find(&userSegments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user segments"})
		return
	}

	var segmentIDs []uint
	for _, userSegment := range userSegments {
		segmentIDs = append(segmentIDs, userSegment.SegmentID)
	}

	var segments []models.Segment
	if err := db.Where("id IN (?)", segmentIDs).Find(&segments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve segments"})
		return
	}

	c.JSON(http.StatusOK, segments)
}
