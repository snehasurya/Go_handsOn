package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Store struct {
	mu   sync.Mutex
	data map[int]User
}

var storeData = Store{
	data: make(map[int]User),
}

func getSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	key, _ := strconv.Atoi(id)
	storeData.mu.Lock()
	defer storeData.mu.Unlock()
	if u, exists := storeData.data[key]; exists {
		json.NewEncoder(w).Encode(u)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	storeData.mu.Lock()
	//fmt.Println(data)
	err := json.NewEncoder(w).Encode(storeData.data)
	storeData.mu.Unlock()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	storeData.mu.Lock()
	defer storeData.mu.Unlock()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	storeData.data[user.Id] = user
	w.WriteHeader(http.StatusCreated)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	key, _ := strconv.Atoi(id)
	storeData.mu.Lock()
	defer storeData.mu.Unlock()
	if _, exists := storeData.data[key]; exists {
		delete(storeData.data, key)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("method %s started since %v", r.Method, startTime)
	})
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			token := r.Header.Get("Authorization")
			if token != "Bearer my token" {
				http.Error(w, "User is unauthorized", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func rateLimiter(next http.Handler) http.Handler {
	rateLimit := rate.NewLimiter(1, 5)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rateLimit.Allow() {
			http.Error(w, "too many request", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	storeData.data[1] = User{Id: 1,
		Name: "sneha"}
	storeData.data[2] = User{Id: 2,
		Name: "Vivaan"}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    "0.0.0.0:8087",
		Handler: mux,
	}
	mux.HandleFunc("POST /users", auth(createUser))
	mux.Handle("GET /users", logger(rateLimiter(http.HandlerFunc(getAllUser))))
	mux.Handle("GET /users/{id}", logger(rateLimiter(http.HandlerFunc(getSingleUser))))
	mux.Handle("DELETE /users/{id}", logger(rateLimiter(auth(http.HandlerFunc(deleteUser)))))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error while starting server %s", err)
	}

}
