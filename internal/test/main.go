package main

/*
import (
	. "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mongodb"
	"log"
	"users/internal/types"
)

func main () {
	session, err := NewMongoDBClientWithCredentials(
		"user-db",
		"jlwnueGGCxpjwlpxTJ6ovBIi5xW3aoCTEignYQNsJTep9e7rHP3AUBs9wADxUduoNZVpfJ4hszbm00lnj8eqcQ==",
		"user-db.mongo.cosmos.azure.com",
	)

	if err != nil {
		log.Fatal(err)
	}

	err = session.Ping()

	if err != nil {
		log.Fatal(err)
	}

	newUser := types.User{
		Name:     "Dennis",
		Email:    "dennis.sentler@haw-hamburg.de",
		Password: "testpw",
	}
	err = session.DB("TEST").C("user").Insert(newUser)

	if err != nil {
		log.Fatal(err)
	}
}*/
