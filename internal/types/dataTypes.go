package types

import "encoding/json"

type FailMessage struct {
	Fault string
}

type User struct {
	Name     string
	Email    string
	Password string
}

func NewUser(binary []byte) (*User, error) {
	var newUser User
	parseErr := json.Unmarshal(binary, &newUser)
	return &newUser, parseErr
}
