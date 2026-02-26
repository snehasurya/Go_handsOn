package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type URLStatus struct {
	Url    string
	Status string
}

func checkUrls(urls []string) []URLStatus {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	responses := make(chan URLStatus, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkUrlStatus(url, client, responses, &wg)
	}
	wg.Wait()
	close(responses)
	allResponses := make([]URLStatus, 0, len(urls))
	for res := range responses {
		allResponses = append(allResponses, res)
	}
	return allResponses
}

func checkUrlStatus(url string, client *http.Client, responses chan<- URLStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	var status URLStatus
	status.Url = url

	resp, err := client.Head(url)

	if err != nil {
		status.Status = fmt.Sprintf("DOWN (Error: %v)", err.Error())
	} else {
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			status.Status = fmt.Sprintf("DOWN (HTTP Status: %s)", resp.Status)
		} else {
			status.Status = fmt.Sprintf("UP (HTTP Status: %s)", resp.Status)
		}
	}
	responses <- status
}
func nomain() {
	urlsToCheck := []string{
		"https://www.google.com",
		"https://www.nonexistentwebsite12345.org", // Should fail (down)
		"https://pkg.go.dev",
		"http://httpbin.org/status/500", // Should be reachable but report an error status
		"https://developer.mozilla.org",
	}

	fmt.Printf("Checking %d URLs concurrently...\n", len(urlsToCheck))

	finalResults := checkUrls(urlsToCheck)

	fmt.Println("\n--- Check Results ---")
	for _, res := range finalResults {
		fmt.Printf("%s: %s\n", res.Url, res.Status)
	}
}
