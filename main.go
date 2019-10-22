package main

import (
	"fmt"
	"users/internal/api"
	"users/internal/db"
)

func main() {
	err := db.Connect("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
	}

	api.Init(8080)
}
