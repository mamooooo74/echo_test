package main

import (
	"fmt"
	"test_api/api/model"
	"test_api/util"
)

func main() {
	users := []model.User{}
	db := util.NewDB()
	defer util.CloseDB(db)
	db.Find(&users)
	fmt.Println(users)
}
