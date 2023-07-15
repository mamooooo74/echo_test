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
}
