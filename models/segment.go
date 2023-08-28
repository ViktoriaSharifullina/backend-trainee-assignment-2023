package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Segment struct {
	ID   uint `gorm:"primary_key"`
	Slug string
}

type UserSegment struct {
	UserID            uint       `gorm:"primary_key"`
	SegmentID         uint       `gorm:"primary_key"`
	ExpiresAt         *time.Time `json:"expires_at"`
	AutoAssignPercent int        `json:"auto_assign_percent"`
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

func CreateUserSegment(db *gorm.DB, userSegment UserSegment, ttl int) error {
	//if ttl > 0 {
	//	expirationTime := time.Now().Add(time.Duration(ttl))
	//	userSegment.ExpiresAt = &expirationTime
	//}

	if err := db.Create(&userSegment).Error; err != nil {
		return err
	}
	return nil
}
