package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"users/internal/services/userauth"
	. "users/internal/types"
)

func Init(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/user", handleRegisterUser).Methods("POST")
	router.HandleFunc("/user/login", handleLogin).Methods("GET")
	router.HandleFunc("/user/logout", handleLogout).Methods("GET")
	router.HandleFunc("/user/auth", handleAuth).Methods("GET")
	router.HandleFunc("/user/{email}", handleGetUser).Methods("GET")
	router.HandleFunc("/user/{email}", handleUpdateUser).Methods("PUT")
	router.HandleFunc("/user/{email}", handleDeleteUser).Methods("DELETE")
	router.HandleFunc("/user/shutdown", handleShutdown)

	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), router))
}

func handleShutdown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("debug", "ANY /shutdown")
	discErr := userauth.Disconnect()
	if discErr == nil {
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("down"))
	} else {
		_ = sendJSONResponse(&w, FailMessage{Fault: discErr.Error()}, http.StatusConflict)
	}
	time.Sleep(2)
	os.Exit(0)
}

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	//debugging
	w.Header().Set("debug", "POST /userService")

	var userData *User
	var parseErr error
	userData, parseErr = getUserFromBody(r.Body)
	if parseErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: parseErr.Error()}, http.StatusBadRequest)
		return
	}

	regErr := userauth.Register(*userData)
	if regErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: regErr.Error()}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	tokenstring := r.Header.Get("Authorization")
	tokenstring = strings.TrimPrefix(tokenstring, "Bearer ")
	jwt, authErr := userauth.Auth(tokenstring)

	if authErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: authErr.Error()}, http.StatusUnauthorized)
		return
	}

	_ = sendJSONResponse(&w, jwt, http.StatusOK)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		_ = sendJSONResponse(&w, FailMessage{Fault: "username and password are not formatted correctly"}, http.StatusUnauthorized)
		return
	}

	jwt, loginErr := userauth.Login(username, password)
	if loginErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: loginErr.Error()}, http.StatusUnauthorized)
		return
	} else {
		w.Header().Set("Authorization", "Bearer "+jwt)
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	tokenstring := r.Header.Get("Authorization")
	tokenstring = strings.TrimPrefix(tokenstring, "Bearer ")
	authErr := userauth.Logout(tokenstring)

	if authErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: authErr.Error()}, http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	tokenstring := r.Header.Get("Authorization")
	tokenstring = strings.TrimPrefix(tokenstring, "Bearer ")
	vars := mux.Vars(r)
	email := vars["email"]
	jwt, authErr := userauth.Auth(tokenstring)

	if authErr != nil || jwt.Email != email {
		_ = sendJSONResponse(&w, FailMessage{Fault: authErr.Error()}, http.StatusUnauthorized)
		return
	}

	var newUserP *User
	newUserP, parseErr := getUserFromBody(r.Body)

	if parseErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: parseErr.Error()}, http.StatusBadRequest)
		return
	}

	updErr := userauth.UpdateUser(email, *newUserP)
	if updErr != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: updErr.Error()}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenstring := r.Header.Get("Authorization")
	tokenstring = strings.TrimPrefix(tokenstring, "Bearer ")
	vars := mux.Vars(r)
	email := vars["email"]
	jwt, authErr := userauth.Auth(tokenstring)

	if authErr != nil || jwt.Email != email {
		_ = sendJSONResponse(&w, FailMessage{Fault: authErr.Error()}, http.StatusUnauthorized)
		return
	}

	err := userauth.DeleteUser(email)

	if err != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: err.Error()}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	tokenstring := r.Header.Get("Authorization")
	tokenstring = strings.TrimPrefix(tokenstring, "Bearer ")
	vars := mux.Vars(r)
	email := vars["email"]
	jwt, authErr := userauth.Auth(tokenstring)

	if authErr != nil || jwt.Email != email {
		_ = sendJSONResponse(&w, FailMessage{Fault: authErr.Error()}, http.StatusUnauthorized)
		return
	}

	user, err := userauth.GetUser(email)

	if err != nil {
		_ = sendJSONResponse(&w, FailMessage{Fault: err.Error()}, http.StatusInternalServerError)
		return
	}

	_ = sendJSONResponse(&w, user, http.StatusOK)
}

func getUserFromBody(body io.ReadCloser) (*User, error) {
	var user *User
	bin, err := ioutil.ReadAll(body)
	if err != nil {
		return user, err
	} else {
		var parseErr error
		user, parseErr = NewUserBin(bin)
		if parseErr != nil {
			return user, parseErr
		}
	}
	return user, nil
}

func sendJSONResponse(wp *http.ResponseWriter, obj interface{}, status int) error {
	w := *wp
	jsonRes, err := json.Marshal(obj)
	if err != nil {
		return err
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(jsonRes)
		return nil
	}
}
