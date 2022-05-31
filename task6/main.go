package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Function which kills goroutine by context calling
func closeWithContext(wg *sync.WaitGroup) {
	// Context with timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	// Start goroutine with context
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Print("close with context\n")
				wg.Done()
				return
			default:
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)
}

// Function which kills goroutine by sending stop signal through the channel
func closeWithSignal(wg *sync.WaitGroup) {
	// Creating data and signal channels
	doneChan := make(chan interface{})
	dataChan := make(chan int)
	// Starting goroutine which listening two channels
	go func(data <-chan int, done <-chan interface{}) {
		for {
			select {
			case num := <-data:
				fmt.Printf("got data: %d\n", num)
			case <-done:
				fmt.Print("close with signal\n")
				wg.Done()
				return
			}
		}
	}(dataChan, doneChan)
	dataChan <- 0
	time.Sleep(1 * time.Second)
	doneChan <- struct{}{}
	defer close(doneChan)
	defer close(dataChan)
}

// Function which kills goroutine by closing data channel
func closeWithClosing(wg *sync.WaitGroup) {
	// Create data channel
	dataChan := make(chan int)
	defer close(dataChan)
	go func(data <-chan int) {
		// Listening channel while it opens
		for _ = range data {

		}
		fmt.Print("close with closing\n")
		wg.Done()
	}(dataChan)
	time.Sleep(3 * time.Second)
}

func main() {
	// Create wait group for all goroutines
	var wg sync.WaitGroup
	// Add three workers to waitGroup
	wg.Add(3)
	// Starts goroutines
	closeWithContext(&wg)
	closeWithSignal(&wg)
	closeWithClosing(&wg)
	// Starts waiting
	wg.Wait()
}
