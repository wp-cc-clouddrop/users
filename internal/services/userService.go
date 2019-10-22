package services

import (
	"errors"
	"users/internal/db"
	. "users/internal/types"
)

func Register(newUser User) error {
	err := db.Create("user", newUser)
	return err
}

func GetUser() (User, error) {
	user := User{"testname", "test@mail", "testpw"}
	return user, nil
}

func UpdateUser(newUser User) error {
	return errors.New("NoUserFound")
}

func DeleteUser(id string) error {
	return errors.New("NoUserFound")
}
