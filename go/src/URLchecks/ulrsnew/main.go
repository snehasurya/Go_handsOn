// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type UrlStatus struct {
	url        string
	statusCode string
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://www.google.com",
		"https://www.abc.com",
		"https://www.nonexistentwebsite12345.org",
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	result := make(chan UrlStatus, 2)
	for _, url := range urls {
		wg.Add(1)
		go urlCheck(url, client, result, &wg)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	output := make([]UrlStatus, 0)
	for r := range result {
		output = append(output, r)
	}
	fmt.Println(output)
}

func urlCheck(urls string, client *http.Client, result chan UrlStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	var res UrlStatus
	res.url = urls
	resp, err := client.Head(urls)

	if err != nil {
		res.statusCode = fmt.Sprintf(" err is %s", err.Error())
		result <- res
		return
	}
	defer resp.Body.Close()
	res.statusCode = resp.Status
	result <- res

}
