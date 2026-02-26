package main

import (
	"encoding/json"
	"net/http"
)

var msg string = "Hello"

type credentials struct {
	id       string
	password string
}

func main() {
	c := credentials{
		id:       "abc",
		password: "dce",
	}
	http.HandleFunc("/messages", getmessage)

	http.ListenAndServe(":8082", nil)
}

func getmessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var cred credentials
	auth := r.Header.Get("Authorization")
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.Write(http.StatusBadRequest, err)
		return
	}

	json.NewEncoder(w).Encode(msg)

}
