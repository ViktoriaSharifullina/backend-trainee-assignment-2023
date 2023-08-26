package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testAvito/models"
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

	c.JSON(http.StatusOK, gin.H{"segment": segment_db.Value})
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

func AddUserToSegment(c *gin.Context) {
	db := models.GetDB()

	segmentSlug := c.Param("segment_slug")
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var segment models.Segment
	if err := db.Where("slug = ?", segmentSlug).First(&segment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userSegment := models.UserSegment{UserID: uint(userID), SegmentID: segment.ID}
	if err := models.CreateUserSegment(db, userSegment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to segment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added to segment successfully"})
}
