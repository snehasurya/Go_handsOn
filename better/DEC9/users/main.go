package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type emailVerify struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

func (e emailVerify) Error() string {
	return fmt.Sprintf("Response code %d and error is %s", e.code, e.msg)
}

var users []User

var nextId int

func main() {

	http.HandleFunc("/users", createUser)
	http.ListenAndServe(":8070", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	if strings.TrimSpace(newUser.Email) == "" {
		emailCheck := emailVerify{
			code: http.StatusBadRequest,
			msg:  "Email field cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(emailCheck)
		log.Printf("Email validation error for user %s", newUser.Name)
		return
	}
	newUser.Id = nextId
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
