package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
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
	//var err error
	//db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=avito_db password=1234 sslmode=disable")
	//if err != nil {
	//	panic("Failed to connect to database")
	//}
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD")))
	if err != nil {
		panic("Failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}
