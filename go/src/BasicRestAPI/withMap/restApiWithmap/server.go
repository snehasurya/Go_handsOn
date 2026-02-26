package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SafeStore struct {
	data []User
	mu   sync.Mutex
}

var store = SafeStore{data: make([]User, 0)}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		store.mu.Lock()
		data := store.data
		store.mu.Unlock()
		fmt.Println(data)
		json.NewEncoder(w).Encode(data)
	case "POST":
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		store.mu.Lock()
		store.data = append(store.data, u)
		store.mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	}
}

func getSingleUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for _, u := range store.data {
		if id == u.Id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			token := r.Header.Get("Authorization")
			fmt.Println(token)
			if token != "Bearer my token" {
				http.Error(w, "Unauthorised: Invalid or missing token", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	}
}
func main() {
	//http.HandleFunc("/users", userHandler)
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}

	mux.HandleFunc("GET /users/{id}", middleware(getSingleUser))
	mux.HandleFunc("GET /users", middleware(userHandler))
	mux.HandleFunc("POST /users", middleware(userHandler))
	users := []User{
		{"a", "sneha"},
		{"b", "pooja"},
	}
	store.data = append(store.data, users...)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error %s", err)
		}

	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("server shuting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Nanosecond)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("closing")
		log.Fatal("server forced to shut down", err)
	}
	log.Println("datebase close")
	log.Println("server exited cleanly")

}
