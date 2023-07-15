package main

import (
	"fmt"
	"test_api/api/model"
	"test_api/util"
)

func main() {
	db := util.NewDB()
	defer fmt.Println("migrate success")
	defer util.CloseDB(db)
	db.AutoMigrate(&model.User{})
}
