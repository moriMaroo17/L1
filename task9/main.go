package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Sends elems from the array to the channel
func sender(wg *sync.WaitGroup, sendChan chan<- int) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, elem := range arr {
		sendChan <- elem
	}
	defer wg.Done()
	defer close(sendChan)
}

// Gets elements from the one channel,
// count their square and resend to other chanel
func converter(wg *sync.WaitGroup, getChan <-chan int, sendChan chan<- int) {
	for num := range getChan {
		sendChan <- num * num
	}
	defer wg.Done()
	defer close(sendChan)
}

// Get elements squares from the channel and send them to stdout
func printer(wg *sync.WaitGroup, getChan <-chan int) {
	for num := range getChan {
		io.WriteString(os.Stdout, fmt.Sprintf("%d\n", num))
	}
	defer wg.Done()
}

func main() {
	// Create wg, and add 3 task to it
	var wg sync.WaitGroup
	wg.Add(3)
	// Create two channels
	originChan := make(chan int)
	squareChan := make(chan int)
	// Run 3 goroutines
	go sender(&wg, originChan)
	go converter(&wg, originChan, squareChan)
	go printer(&wg, squareChan)
	// Wait for all goroutines done
	wg.Wait()
}
