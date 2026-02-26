package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Employee struct {
	Id    int    `json:"Id,omitempty"`
	Name  string `json:"Name"`
	Email string `json:"email"`
}

var emps []Employee

var nextId int = 4

func main() {

	emps = append(emps, Employee{Id: 1, Name: "Sneha", Email: "abc@gmail.com"})
	emps = append(emps, Employee{Id: 2, Name: "asdf", Email: "asdf@gmail.com"})
	emps = append(emps, Employee{Id: 3, Name: "qwerty", Email: "qwerty@gmail.com"})

	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees/create", createEmployees)

	log.Fatal(http.ListenAndServe(":8089", nil))
	go func() {
		time.Sleep(500 * time.Millisecond)
		clientPostRequest()
		time.Sleep(500 * time.Millisecond)
	}()

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Retry-After", fmt.Sprintf("%d", 1*time.Second))
	json.NewEncoder(w).Encode(emps)
}

func createEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEmp Employee
	err := json.NewDecoder(r.Body).Decode(&newEmp)
	if err != nil {
		http.Error(w, "Invalid reuqest", http.StatusBadRequest)
		return
	}
	newEmp.Id = nextId
	nextId++
	emps = append(emps, newEmp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmp)

}

func clientPostRequest() {
	postEmp := Employee{
		Name:  "zxc",
		Email: "zxc@gmail.com",
	}
	jsonPayload, err := json.Marshal(postEmp)
	if err != nil {
		fmt.Println("error in marshelling", err)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8089/employees/create", bytes.NewBuffer(jsonPayload))

	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	resPost, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error posting item: %v\n", err)
		return
	}
	defer resPost.Body.Close()

}
