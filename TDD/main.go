package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/employee", getEmployee)
	http.ListenAndServe(":8181", nil)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	employee := []string{
		"Hello", "Sneha",
	}
	fmt.Println(employee)
	json.NewEncoder(w).Encode(employee)
	//w.Write([]byte(`{"status":"OK", "messages":[]}`))
	//w.WriteHeader(http.StatusOK)
}
