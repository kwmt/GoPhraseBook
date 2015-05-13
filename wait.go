package main

import (
	"sync"
	"os"
	"strings"
	"fmt"
)

func main() {
	var w sync.WaitGroup

	for _, v := range os.Args {
		// Increment the WaitGroup counter.
		w.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(str string) {
			// Fetch the URL.
			fmt.Printf("%s\n", strings.ToUpper(str))
			
			// Decrement the counter.
			w.Done()
		}(v)
	}
	// Wait for all HTTP fetches to complete.
	w.Wait()
}
