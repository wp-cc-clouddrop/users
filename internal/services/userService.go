package services

import (
	"errors"
	"log"
	. "users/internal/mongoDB/adapter"
	. "users/internal/mongoDB/adapter/azureCosmos"
	. "users/internal/types"
)

var (
	mongoDB      MongoDBCloud
	myCollection string
)

func init() {
	mongoDB = &AzureCosmos{}
	myCollection = "user"
	connectErr := mongoDB.Connect()
	if connectErr != nil {
		log.Fatal(connectErr)
	}
}

func Disconnect() error {
	return mongoDB.Disconnect()
}

func Register(newUser User) error {
	err := mongoDB.Insert(myCollection, newUser)
	return err
}

func GetUser(email string) (User, error) {
	bin, getErr := mongoDB.Get(myCollection, email)
	if getErr != nil { //not found
		return User{}, getErr
	}

	user, parseErr := NewUserFromDB(bin) //parse data to User Struct
	return *user, parseErr
}

func UpdateUser(email string, newUser User) error {
	updateErr := mongoDB.Update(myCollection, email, newUser)
	return updateErr
}

func DeleteUser(id string) error {
	return errors.New("NoUserFound")
}
