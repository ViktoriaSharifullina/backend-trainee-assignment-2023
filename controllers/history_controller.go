package controllers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testAvito/models"
	"time"
)

func GenerateSegmentHistoryReport(c *gin.Context) {
	db := models.GetDB()

	year := c.Query("year")
	month := c.Query("month")

	reportFolderPath := "reports"

	// Получение записей о попаданиях/выбываниях пользователей из сегментов за указанный период
	var historyRecords []models.UserSegmentHistory
	if err := db.Where("EXTRACT(YEAR FROM date) = ? AND EXTRACT(MONTH FROM date) = ?", year, month).Find(&historyRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve segment history"})
		return
	}

	// Создание CSV файла
	filename := "segment_history_" + year + "_" + month + ".csv"
	file, err := os.Create(filepath.Join(reportFolderPath, filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create report file"})
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close report file"})
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовка в CSV файл
	header := []string{"User ID", "Segment", "Operation", "Date and Time"}
	if err := writer.Write(header); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write header"})
		return
	}

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
		if err := writer.Write(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write data"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Segment history report generated successfully", "link": filename})
}
