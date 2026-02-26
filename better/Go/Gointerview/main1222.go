package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var employees []Employee

func main222() {

	employees = append(employees, Employee{Id: 1, Name: "abc", Email: "abc@gmail.com"})
	employees = append(employees, Employee{Id: 2, Name: "d", Email: "abc@gmail.com"})
	employees = append(employees, Employee{Id: 3, Name: "q", Email: "abc@gmail.com"})
	employees = append(employees, Employee{Id: 4, Name: "z", Email: "abc@gmail.com"})

	http.HandleFunc("/employees/", GetEmployee)
	http.HandleFunc("/employees/{Id}", getSingleEmployee)
	http.HandleFunc("/employeesPost/", createEmployee)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleRequestForEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if (r.Method) == "POST" {
		createEmployee(w, r)
	}
}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
	// for _, emp := range employees {
	// 	json.NewEncoder(w).Encode(emp)
	// }
	//http.Error(w, "Employee not found", http.StatusNotFound)
}
func getSingleEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/employees/")
	for _, emp := range employees {
		if id == strconv.Itoa(emp.Id) {
			json.NewEncoder(w).Encode(emp)
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	employees = append(employees, newEmployee)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmployee)
}
