package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

type Item struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Monitor", Price: 350.00},
	{ID: 2, Name: "Keyboard", Price: 75.00},
}
var nextID = 3 // Simple ID counter

// itemHandler processes requests to /api/items
func itemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Always respond with JSON

	switch r.Method {
	case "GET":
		// --- Server: Handling GET ---
		// Serialize the Go slice of Items into JSON
		json.NewEncoder(w).Encode(items)
		// Default status is 200 OK, no explicit call needed

	case "POST":
		// --- Server: Handling POST ---
		var newItem Item

		// Deserialize the request body (JSON) into a Go struct
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Apply Business Logic (e.g., assign ID and save to 'database')
		newItem.ID = nextID
		nextID++
		items = append(items, newItem)

		// Respond with the correct status code and the created item
		w.WriteHeader(http.StatusCreated) // Set status to 201 Created
		json.NewEncoder(w).Encode(newItem)

	default:
		// Handle unsupported methods
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
	}
}

func main() {
	// 1. Server: Routing - Map the path to the handler function

	mux := http.NewServeMux()
	// 2. Server: Listening - Start the server on a specific port
	port := 8085
	mux.HandleFunc("/api/items", middleware(itemHandler))

	fmt.Printf("Server: Starting API server on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		fmt.Printf("Server: Failed to start server: %v\n", err)
	}
}

func middleware(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	limiter := rate.NewLimiter(2, 5)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if !limiter.Allow() {
		// 	http.Error(w, "rate limit exceeds", http.StatusTooManyRequests)
		// 	return
		// }
		reservation := limiter.Reserve()
		delay := reservation.Delay()

		if delay > 0 {
			if !reservation.OK() {
				http.Error(w, "rate limit exceeds", http.StatusTooManyRequests)
				return
			}
			waitTime := delay + 50*time.Millisecond
			retryAfterSeconds := math.Ceil(waitTime.Seconds())
			retryAfter := int(retryAfterSeconds)
			//fmt.Printf("wait time : %d retryretryAfter : %d delay : %d \n", waitTime, retryAfter, delay)
			w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
			//fmt.Printf("writer : %+v", w)
			http.Error(w, "rate limit exceeds", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	})

}
