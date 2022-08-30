package main

import (
	"log"
	"fmt"
	_"log"
	"net/http"
)

var users = map[string]string{
	"test": "secret",
}

func isAuthorised(username, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}

	return password == pass
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "welcome to golang world!"}`))
	return
}

func main() {
	http.HandleFunc("/", greeting)
	fmt.Println("Starting Server at port :8180")
	log.Fatal(http.ListenAndServe(":8180", nil))
}