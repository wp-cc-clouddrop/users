package main

import (
	"log"
	"net/http"
)

func user(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.WriteHeader(404)
		w.Header().Set("debug", "GET /user must be used with identifier")
	case "POST":
		w.WriteHeader(204)
		w.Header().Set("debug", "POST /user")
	case "PUT":
		w.WriteHeader(404)
		w.Header().Set("debug", "PUT /user must be used with identifier")
	case "DELETE":
		w.WriteHeader(404)
		w.Header().Set("debug", "DELETE /user must be used with identifier")
	default:
		w.WriteHeader(404)
		w.Header().Set("debug", "No http verb found")
	}
}

func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}
}
func main() {
	http.HandleFunc("/user", user)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/logout", logout)
	http.HandleFunc("/user/auth", auth)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
