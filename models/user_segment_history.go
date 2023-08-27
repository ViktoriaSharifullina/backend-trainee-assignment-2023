package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserSegmentHistory struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	SegmentID uint      `json:"segment_id"`
	Segment   Segment   `json:"segment" gorm:"foreignKey:SegmentID"`
	Operation string    `json:"operation"` // "add" or "remove"
	Date      time.Time `json:"date"`
}

func CreateHistory(db *gorm.DB, history UserSegmentHistory) error {
	if err := db.Create(&history).Error; err != nil {
		return err
	}
	return nil
}
