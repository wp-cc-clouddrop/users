package userauth

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	. "users/internal/types"
	. "users/internal/userDB/adapter"
	//. "users/internal/userDB/adapter/azureCosmos"
	. "users/internal/userDB/adapter/gcpFirestore"
	"users/internal/utils/jwt"
	"users/internal/utils/pwSecurity"
)

var (
	userDB         UserDBCloud
	userCollection string
	authCollection string
)

func init() {
	cloud := os.Getenv("CLUSTER_ENV")
	if cloud == "gcp" {
		userDB = &GCPFirestore{}
	} else if cloud == "azure" {

	} else {
		println("no valid env var CLUSTER_ENV found, was: " + cloud)
		log.Fatal("no valid env var CLUSTER_ENV found, was: " + cloud)
		os.Exit(1)
	}

	userCollection = "user"
	authCollection = "auth"
	connectErr := userDB.Connect()
	if connectErr != nil {
		println(connectErr.Error())
		log.Fatal(connectErr)
		os.Exit(1)
	}
}

func Disconnect() error {
	return userDB.Disconnect()
}

func Register(newUser User) error {
	newUser.Password = pwSecurity.HashAndSalt(newUser.Password)
	err := userDB.Insert(userCollection, newUser)
	return err
}

func GetUser(email string) (User, error) {
	userMap, getErr := userDB.Get(userCollection, email)
	if getErr != nil { //not found
		return User{}, getErr
	}
	var user User
	mapErr := mapstructure.Decode(userMap, &user)
	if mapErr != nil {
		return User{}, errors.New("jwt, mapping error")
	}
	//user, parseErr := NewUserFromDB(bin) //parse data to User Struct
	return user, nil
}

func UpdateUser(email string, newUser User) error {
	newUser.Password = pwSecurity.HashAndSalt(newUser.Password)
	updateErr := userDB.Update(userCollection, email, newUser)
	return updateErr
}

func DeleteUser(email string) error {
	deleteErr := userDB.Delete(userCollection, email)
	if deleteErr == nil {
		deleteErr = userDB.Delete(authCollection, email)
	}
	return deleteErr
}

func Login(email string, password string) (string, error) {
	userMap, lookupErr := userDB.Get(userCollection, email)
	if lookupErr != nil {
		return "", lookupErr
	}
	//user, _ := NewUserFromDB(userBin) // with hashed password
	var user User
	mapErr := mapstructure.Decode(userMap, &user)
	if mapErr != nil {
		return "", errors.New("jwt, mapping error")
	}
	pass := pwSecurity.ComparePasswords(user.Password, password)

	if pass {
		token, lookupErr := jwt.GenerateJWT(email, true)
		jwt := JWT{
			JWT:   token,
			Email: email,
		}
		insErr := userDB.Insert(authCollection, jwt)
		if insErr != nil {
			updErr := userDB.Update(authCollection, email, jwt)
			return token, updErr
		} else {
			return token, lookupErr
		}
	} else {
		return "", errors.New("login failed, password incorect")
	}
}

func Logout(tokenstring string) error {
	jwtMap, findErr := userDB.Find(authCollection, "jwt", tokenstring)
	if findErr != nil {
		return errors.New("jwt, session-key is not valid")
	}
	var jwt JWT
	mapErr := mapstructure.Decode(jwtMap, &jwt)
	if mapErr != nil {
		return errors.New("jwt, mapping error")
	}

	//jwt, _ := NewJWTFromDB(jwtbin)
	delErr := userDB.Delete(authCollection, jwt.Email)
	return delErr
}

func Auth(tokenstring string) (JWT, error) {
	jwtMap, findErr := userDB.Find(authCollection, "jwt", tokenstring)
	if findErr != nil {
		return JWT{}, errors.New("jwt, session-key is not valid")
	}
	var jwt JWT
	mapErr := mapstructure.Decode(jwtMap, &jwt)
	if mapErr != nil {
		return JWT{}, errors.New("jwt, mapping error")
	}
	//jwt, parseErr := NewJWTFromDB(jwtbin)
	return jwt, nil
}
