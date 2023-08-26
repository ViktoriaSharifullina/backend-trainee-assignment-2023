package models

import (
	"github.com/jinzhu/gorm"
)

type Segment struct {
	ID   uint `gorm:"primary_key"`
	Slug string
}

type UserSegment struct {
	UserID    uint `gorm:"primary_key"`
	SegmentID uint `gorm:"primary_key"`
}

func CreateSegment(db *gorm.DB, segment Segment) error {
	if err := db.Create(&segment).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserSegment(db *gorm.DB, userSegment UserSegment) error {
	if err := db.Create(&userSegment).Error; err != nil {
		return err
	}
	return nil
}
