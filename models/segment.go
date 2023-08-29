package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Segment struct {
	ID                uint `gorm:"primary_key"`
	Slug              string
	AutoAssignPercent int `json:"auto_assign_percent"`
}

type UserSegment struct {
	UserID    uint       `gorm:"primary_key"`
	SegmentID uint       `gorm:"primary_key"`
	ExpiresAt *time.Time `json:"expires_at"`
}

func CreateSegment(db *gorm.DB, segment Segment) error {
	if err := db.Create(&segment).Error; err != nil {
		return err
	}
	return nil
}

func GetSegmentByID(db *gorm.DB, segmentID uint) (*Segment, error) {
	var segment Segment
	if err := db.First(&segment, segmentID).Error; err != nil {
		return nil, err
	}
	return &segment, nil
}

func CreateUserSegment(db *gorm.DB, userSegment UserSegment) error {
	if err := db.Create(&userSegment).Error; err != nil {
		return err
	}
	return nil
}
