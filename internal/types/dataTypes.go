package types

import (
	"encoding/json"
	"errors"
)

type FailMessage struct {
	Fault string
}

type MongoAccess struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type DBEntry interface {
}

type User struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"_id" json:"email"`
	Password string `bson:"password" json:"password"`
}

type dbUser struct { //needed ugly workaround, at least its private
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"_id" json:"_id"`
	Password string `bson:"password" json:"password"`
}

//needed ugly workaround
func NewUserFromDB(binary []byte) (*User, error) {
	var newUser User
	var newDBUser dbUser
	parseErr := json.Unmarshal(binary, &newDBUser)
	newUser = User(newDBUser)
	if parseErr == nil {
		if (len(newUser.Email) == 0) || (len(newUser.Password) == 0) {
			parseErr = errors.New("provided email and password must be NOT empty")
		}
	}
	return &newUser, parseErr
}

func NewUserBin(binary []byte) (*User, error) {
	var newUser User
	parseErr := json.Unmarshal(binary, &newUser)

	if parseErr == nil {
		if (len(newUser.Email) == 0) || (len(newUser.Password) == 0) {
			parseErr = errors.New("provided email and password must be NOT empty")
		}
	}
	return &newUser, parseErr
}
