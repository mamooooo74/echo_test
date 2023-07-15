package util

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	url := "root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("database connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDb, _ := db.DB()
	if err := sqlDb.Close(); err != nil {
		log.Fatalln(err)
	}
}
