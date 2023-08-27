package controllers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"testAvito/models"
	"time"
)

func GenerateSegmentHistoryReport(c *gin.Context) {
	db := models.GetDB()

	year := c.Query("year")
	month := c.Query("month")

	// Получение записей о попаданиях/выбываниях пользователей из сегментов за указанный период
	var historyRecords []models.UserSegmentHistory
	if err := db.Where("YEAR(date) = ? AND MONTH(date) = ?", year, month).Find(&historyRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve segment history"})
		return
	}

	// Создание CSV файла
	filename := "segment_history_" + year + "_" + month + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create CSV file"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовка в CSV файл
	header := []string{"User ID", "Segment", "Operation", "Date and Time"}
	writer.Write(header)

	// Запись данных в CSV файл
	for _, record := range historyRecords {
		segment, err := models.GetSegmentByID(db, record.SegmentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch segment"})
			return
		}

		row := []string{
			strconv.Itoa(int(record.UserID)),
			segment.Slug,
			record.Operation,
			record.Date.Format(time.RFC3339),
		}
		writer.Write(row)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment history report generated successfully", "link": filename})
}
