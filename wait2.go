package main

import (
	"sync"
	"net/http"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Fetch the URL.
			http.Get(url)
			// Decrement the counter.
			wg.Done()
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
