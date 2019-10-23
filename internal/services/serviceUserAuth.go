package services

import (
	"log"
	. "users/internal/mongoDB/adapter"
	. "users/internal/mongoDB/adapter/azureCosmos"
	. "users/internal/types"
	"users/internal/utils"
)

var (
	mongoDB        MongoDBCloud
	userCollection string
	authCollection string
)

func init() {
	mongoDB = &AzureCosmos{}
	userCollection = "user"
	authCollection = "auth"
	connectErr := mongoDB.Connect()
	if connectErr != nil {
		log.Fatal(connectErr)
	}
}

func Disconnect() error {
	return mongoDB.Disconnect()
}

func Register(newUser User) error {
	newUser.Password = utils.HashAndSalt(newUser.Password)
	err := mongoDB.Insert(userCollection, newUser)
	return err
}

func GetUser(email string) (User, error) {
	bin, getErr := mongoDB.Get(userCollection, email)
	if getErr != nil { //not found
		return User{}, getErr
	}

	user, parseErr := NewUserFromDB(bin) //parse data to User Struct
	return *user, parseErr
}

func UpdateUser(email string, newUser User) error {
	updateErr := mongoDB.Update(userCollection, email, newUser)
	return updateErr
}

func DeleteUser(email string) error {
	deleteErr := mongoDB.Delete(userCollection, email)
	return deleteErr
}

func Login(username string, password string) (string, error) {
	return "nil", nil
}

func Logout() {

}

func Auth() {

}
