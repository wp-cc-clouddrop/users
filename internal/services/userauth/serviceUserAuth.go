package userauth

import (
	"errors"
	"log"
	. "users/internal/mongoDB/adapter"
	. "users/internal/mongoDB/adapter/azureCosmos"
	. "users/internal/types"
	"users/internal/utils/jwt"
	"users/internal/utils/pwSecurity"
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
	newUser.Password = pwSecurity.HashAndSalt(newUser.Password)
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
	newUser.Password = pwSecurity.HashAndSalt(newUser.Password)
	updateErr := mongoDB.Update(userCollection, email, newUser)
	return updateErr
}

func DeleteUser(email string) error {
	deleteErr := mongoDB.Delete(userCollection, email)
	if deleteErr == nil {
		deleteErr = mongoDB.Delete(authCollection, email)
	}
	return deleteErr
}

func Login(email string, password string) (string, error) {
	userBin, lookupErr := mongoDB.Get(userCollection, email)
	if lookupErr != nil {
		return "", lookupErr
	}
	user, _ := NewUserFromDB(userBin) // with hashed password
	pass := pwSecurity.ComparePasswords(user.Password, password)

	if pass {
		token, lookupErr := jwt.GenerateJWT(email, true)
		jwt := JWT{
			JWT:   token,
			Email: email,
		}
		insErr := mongoDB.Insert(authCollection, jwt)
		if insErr != nil {
			updErr := mongoDB.Update(authCollection, email, jwt)
			return token, updErr
		} else {
			return token, lookupErr
		}
	} else {
		return "", errors.New("login failed, password incorect")
	}
}

func Logout(tokenstring string) error {
	jwtbin, findErr := mongoDB.Find(authCollection, "jwt", tokenstring)
	if findErr != nil {
		return errors.New("jwt, session-key is not valid")
	}
	jwt, _ := NewJWTFromDB(jwtbin)
	delErr := mongoDB.Delete(authCollection, jwt.Email)
	return delErr
}

func Auth(tokenstring string) (JWT, error) {
	jwtbin, findErr := mongoDB.Find(authCollection, "jwt", tokenstring)
	if findErr != nil {
		return JWT{}, errors.New("jwt, session-key is not valid")
	}
	jwt, parseErr := NewJWTFromDB(jwtbin)
	return *jwt, parseErr
}
