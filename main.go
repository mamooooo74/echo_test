package main

import (
	"test_api/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
