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

func GetUser() (User, error) {
	user := User{"testname", "test@mail", "testpw"}
	return user, nil
}

func UpdateUser(email string, newUser User) error {
	updateErr := mongoDB.Update(myCollection, email, newUser)
	return updateErr
}

func DeleteUser(id string) error {
	return errors.New("NoUserFound")
}
