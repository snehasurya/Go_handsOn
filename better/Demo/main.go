package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func mainno() {
	router := mux.NewRouter()
	router.HandleFunc("/about", GetKeys).Methods("GET")
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8010",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Println("Server starting on:", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
