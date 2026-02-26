package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type UrlStatuses struct {
	url    string
	status string
}

func UrlChecks(urls []string) []UrlStatuses {
	result := make([]UrlStatuses, 0, len(urls))
	var wg sync.WaitGroup
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	statusChan := make(chan UrlStatuses, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go CheckUrlStatuses(url, client, statusChan, &wg)
	}
	wg.Wait()
	close(statusChan)

	for status := range statusChan {
		result = append(result, status)
	}
	return result
}

func CheckUrlStatuses(url string, client *http.Client, statusChan chan UrlStatuses, wg *sync.WaitGroup) {
	defer wg.Done()
	var urlStatus UrlStatuses
	urlStatus.url = url
	resp, err := client.Head(url)
	if err != nil {
		urlStatus.status = fmt.Sprintf("the site is down %v", err.Error())
		statusChan <- urlStatus
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		urlStatus.status = fmt.Sprint("Down ", resp.Status)
	} else {
		urlStatus.status = fmt.Sprint("UP ", resp.Status)
	}

	statusChan <- urlStatus
}
func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.nonexistentwebsite12345.org", // Should fail (down)
		"https://pkg.go.dev",
		"http://httpbin.org/status/500", // Should be reachable but report an error status
		"https://developer.mozilla.org",
	}
	result := UrlChecks(urls)
	for _, re := range result {
		fmt.Println(re)
	}
}
