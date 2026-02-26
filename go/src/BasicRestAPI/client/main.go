package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Define the structure for data exchange
type Item struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	apiURL := "http://localhost:8085/api/items" // The server's endpoint
	//getItems(apiURL)
	for i := 0; i < 10; i++ {
		getItems(apiURL)
	}
	//postItem(apiURL)
}

func getItems(apiURL string) {
	maxRetries := 5
	retryCount := 0

	for retryCount < maxRetries {
		// --- 1. Client: GET Request (Fetch data) ---
		fmt.Println("Client: Sending GET request to fetch items...")
		resp, err := http.Get(apiURL)
		if err != nil {
			fmt.Printf("Error fetching items: %v\n", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfterInterval := resp.Header.Get("Retry-After")
			if retryAfterInterval == "" {
				fmt.Printf("GET Response Status: %s\n", resp.Status)
				time.Sleep(1 * time.Second)
			} else {
				waitInterval, parseErr := strconv.Atoi(retryAfterInterval)
				fmt.Printf("waitInterval : %d\n", waitInterval)
				if parseErr != nil {
					fmt.Printf("GET Response Status: %s\n", resp.Status)
					time.Sleep(1 * time.Second)
				} else {
					sleepTime := time.Duration(waitInterval) * time.Second
					fmt.Println("going to sleep for ", sleepTime)
					time.Sleep(sleepTime)
				}
			}
			retryCount++
			continue
			//fmt.Printf("Response from get for Retry-After = %+v \n", resp.Header.Get("Retry-After"))
		} else {
			fmt.Printf("GET Response Status: %s\n", resp.Status)
			io.Copy(io.Discard, resp.Body)
			break
		}
	}
	if retryCount == maxRetries {
		fmt.Printf("client request reached to max retry limit %d\n", maxRetries)
	}
}

func postItem(apiURL string) {
	// Read and print the response body (typically a list of items)
	// Example:

	// --- 2. Client: POST Request (Create data) ---
	newItem := Item{
		Name:  "Laptop",
		Price: 1200.50,
	}

	// Marshal the Go struct into a JSON payload
	jsonPayload, err := json.Marshal(newItem)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println("\nClient: Sending POST request to create an item...")

	// Create the HTTP request object
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}
	postResp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error posting item: %v\n", err)
		return
	}
	defer postResp.Body.Close()

	fmt.Printf("Client: POST Response Status: %s\n", postResp.Status)

	// Example of reading and parsing the server's response (e.g., the created item with an ID)
	body, _ := io.ReadAll(postResp.Body)
	var createdItem Item
	if postResp.StatusCode == http.StatusCreated {
		json.Unmarshal(body, &createdItem)
		fmt.Printf("Client: Successfully created item with ID: %d\n", createdItem.ID)
	}
}
