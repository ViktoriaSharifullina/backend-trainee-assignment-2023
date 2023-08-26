package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
}

func CreateUser(db *gorm.DB, user User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
