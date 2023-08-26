package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=avito_db password=1234 sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	return db
}

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=avito_db password=1234 sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}
