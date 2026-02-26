package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var emps []Employee

func main() {
	emps = append(emps, Employee{Id: 1, Name: "Sneha", Email: "abc@gmail"})
	emps = append(emps, Employee{Id: 2, Name: "Vivaan", Email: "abc@gmail"})
	emps = append(emps, Employee{Id: 3, Name: "viv", Email: "abc@gmail"})

	http.HandleFunc("/emps/", getAllEmps)
	http.HandleFunc("/emps/{id}", getSingleEmps)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getAllEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(emps)
}

func getSingleEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := strings.Trim(r.URL.Path, "/emps")
	fmt.Println("id ", id)
	for _, emp := range emps {
		eId := strconv.Itoa(emp.Id)
		if id == eId {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	http.Error(w, "Emp not found", http.StatusNotFound)
}
