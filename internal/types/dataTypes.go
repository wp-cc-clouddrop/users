package types

import (
	"encoding/json"
	"errors"
)

type FailMessage struct {
	Fault string
}

type DBAccess struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type dbJWT struct {
	JWT   string `bson:"jwt" json:"jwt"`
	Email string `bson:"_id" json:"_id"`
}

type UserDataI interface {
	Id() string
}

type JWT struct {
	JWT   string `bson:"jwt" json:"jwt" firestore:"jwt"`
	Email string `bson:"_id" json:"email" firestore:"email"`
}

func (j JWT) Id() string {
	return j.Email
}

func (u User) Id() string {
	return u.Email
}

type User struct {
	Name     string `bson:"name" json:"name" firestore:"name"`
	Email    string `bson:"_id" json:"email" firestore:"email"`
	Password string `bson:"password" json:"password" firestore:"password"`
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

//needed ugly workaround
func NewJWTFromDB(binary []byte) (*JWT, error) {
	var newJWT JWT
	var newDBJWT dbJWT
	parseErr := json.Unmarshal(binary, &newDBJWT)
	newJWT = JWT(newDBJWT)
	return &newJWT, parseErr
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
