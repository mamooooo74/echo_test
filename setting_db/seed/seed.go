package main

import (
	"log"
	"test_api/api/model"
	"test_api/util"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	db := util.NewDB()
	defer util.CloseDB(db)
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	password := string(hash)
	db.Create([]model.User{
		{Email: "test1@example.com", Password: password},
		{Email: "test2@example.com", Password: password},
		{Email: "test3@example.comd", Password: password},
		{Email: "test4@example.comd", Password: password},
	})
	db.Create([]model.Task{
		{Title: "test1"},
		{Title: "test2"},
		{Title: "test3"},
		{Title: "test4"},
		{Title: "test5"},
		{Title: "test6"},
		{Title: "test7"},
	})
}
