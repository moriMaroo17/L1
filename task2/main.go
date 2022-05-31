package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Create wait group for all goroutines
var wg sync.WaitGroup

func main() {
	arr := [5]uint{2, 4, 6, 8, 10}
	// Go around the array and start goroutine for each element
	for _, elem := range arr {
		// Add one to wait group
		wg.Add(1)
		go func(num uint) {
			io.WriteString(os.Stdout, fmt.Sprintf("%d\n", num*num))
			// Remove one from wait group
			defer wg.Done()
		}(elem)
	}
	// Start waiting
	wg.Wait()
}
