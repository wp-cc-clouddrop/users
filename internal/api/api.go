package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"users/internal/services"
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

	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), router))
}

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	//debugging
	w.Header().Set("debug", "POST /user")

	userDataBin, readErr := ioutil.ReadAll(r.Body)
	var userData *User
	if readErr == nil {
		var parseErr error
		userData, parseErr = NewUser(userDataBin)

		if parseErr != nil {
			w.Header().Set("Content-Type", "application/json")
			jsonResponse, _ := json.Marshal(&FailMessage{Fault: parseErr.Error()})
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(jsonResponse)
		}
	}

	regErr := services.Register(*userData)

	if regErr == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonResponse, _ := json.Marshal(&FailMessage{Fault: regErr.Error()})
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(jsonResponse)
	}
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {

}
