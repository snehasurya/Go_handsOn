// Implement the handler functions GetKeys, GetKey, CreateKey, and DeleteKey. Use
// in-memory storage; the keys do not need to be persisted across restart. The key
// structure is given. How you manage the storage is up to you.
//
// The get methods should only return keys that are unexpired. If the current
// time is after the key's expiration, it should not be returned in the
// response.
//
// All handler methods should return json data.
//
// Example requests and responses:
//
// # Get all unexpired keys.
//
//	curl --verbose --request GET \
//			--url 'http://localhost:8080/keys'
//
// -> 200
//
//	[{"id":"<uuid>","expires":"<iso date>"},{"id":"<uuid>","expires":"<iso date>"}]
//
// # Get a specific key. If the key is expired, the response should be Not Found/404.
//
//	curl --verbose --request GET \
//			--url 'http://localhost:8080/keys/<uuid>'
//
// -> 200
//
//	{"id":"<uuid>","expires":"<iso date>"}
//
// -> 404
//
// # Create a key with no expiration. The expiration should be one hour from
// # when the key is created.
//
//	curl --verbose --request POST \
//			--url 'http://localhost:8080/keys' \
//			--header 'Content-type: application/json' \
//			--data '{}'
//
// -> 204
//
//	{"id":"<uuid>","expires":"<iso date now + 1 hour>"}
//
// # Create a key with explicit expiration. An expires date in the past is accepted.
//
//	curl --verbose --request POST \
//			--url 'http://localhost:8080/keys' \
//			--header 'Content-type: application/json' \
//			--data '{"expires":"<iso date>"}'
//
// -> 204
//
//	{"id":"<uuid>","expires":"<iso date>"}
//
// # Delete a specific key.
//
//	curl --verbose --request DELETE \
//			--url 'http://localhost:8080/keys/<uuid>'
//
// -> 204
// -> 404
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Key struct {
	Id      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

var Keys []Key
var nextKey = 1
var now = time.Now()

func main() {
	Keys = []Key{
		{
			Id:      "1",
			Expires: now.Add(time.Hour * 24 * 7), // Expires 1 week from now (Active)
		},
		{
			Id:      "2",
			Expires: now.Add(time.Minute * 30), // Expires 30 minutes from now (Active)
		},
		{
			Id:      "1",
			Expires: now.Add(time.Minute * -5), // Expired 5 minutes ago (Expired)
		},
		{
			Id:      "2",
			Expires: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC), // Long expired
		},
	}
	//Keys = make([]Key, 0, 10)
	router := mux.NewRouter()
	sub := router.PathPrefix("/keys").Subrouter()
	sub.HandleFunc("", GetKeys).
		Methods("GET")
	sub.HandleFunc("/{key_id}", GetKey).
		Methods("GET")
	sub.HandleFunc("/key", CreateKey).
		Methods("POST").
		HeadersRegexp("Content-Type", "application/json")
	sub.HandleFunc("/{key_id}", DeleteKey).
		Methods("DELETE")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8010",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Get all keys that are not expired. This should return a JSON array of keys.
func GetKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	for _, key := range Keys {
		if time.Now().Compare(key.Expires) == -1 {
			json.NewEncoder(w).Encode(key)
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("GetKeys")
}

// Get a key by id. If the key is expired, return 404. This should return a JSON object.
func GetKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	for _, key := range Keys {
		if key.Id == strings.Trim(r.URL.Path, "/keys") {
			if time.Now().Compare(key.Expires) == -1 {
				json.NewEncoder(w).Encode(key)
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("GetKey")
}

// Create a new key with the specified expiration. If no expiration is provided,
// use the default of 1 hour. Return the new key as the response.
// Example: {"expires": "2019-01-01T12:00:00Z"}
func CreateKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newKey Key
	err := json.NewDecoder(r.Body).Decode(&newKey)
	if err != nil && err.Error() != "EOF" {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	if newKey.Expires.IsZero() {
		newKey.Expires = time.Now().Add(time.Hour * 1)
	}
	newKey.Id = strconv.Itoa(nextKey)
	Keys = append(Keys, newKey)
	fmt.Println(Keys)
	nextKey++
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newKey); err != nil {
		log.Println("Error encoding response:", err)
	}
	log.Printf("CreateKey: Successfully created key ID: %s", newKey.Id)
	//log.Println("CreateKey")
}

// Delete the specified key. If the key is expired, return 404.
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	index := 0
	for _, key := range Keys {
		if key.Id == strings.Trim(r.URL.Path, "/") {
			if time.Now().Compare(key.Expires) == 1 {
				Keys = append(Keys[:index+1], Keys[index+2:]...)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			index++
		}
	}
	w.WriteHeader(http.StatusNoContent)
	log.Println("DeleteKey")
}
