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

type User struct {
	Name     string `json:"name"`
	Email    string `bson:"_id" json:"email"`
	Password string `json:"password"`
}

func NewUser(binary []byte) (*User, error) {
	var newUser User
	parseErr := json.Unmarshal(binary, &newUser)

	if parseErr == nil {
		if (len(newUser.Email) == 0) || (len(newUser.Password) == 0) {
			parseErr = errors.New("provided email and password must be NOT empty")
		}
	}
	return &newUser, parseErr
}
