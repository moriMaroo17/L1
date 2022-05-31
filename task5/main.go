package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Sender function. It get context, wait group link and channel for sending data
func sender(ctx context.Context, wg *sync.WaitGroup, sendChan chan<- any) {
	defer wg.Done()
	for {
		select {
		// If context not done yet, send data to channel
		case <-ctx.Done():
			return
		default:
			sendChan <- time.Now()
		}
	}
}

// Receiver function. It get context, wait group link and channel for get data
func receiver(ctx context.Context, wg *sync.WaitGroup, receiveChan <-chan any) {
	defer wg.Done()
	for {
		select {
		// Listeing channel while context not done
		case data := <-receiveChan:
			fmt.Printf("get data: %v\n", data)
		case <-ctx.Done():
			return
		}
	}
}

// Create main wg
var wg sync.WaitGroup

func main() {
	// Parse second console arg for get timeout in seconds
	if len(os.Args) < 2 {
		log.Fatal("Enter timeout in seconds after program name.")
	}
	timeoutAsStr := os.Args[1]
	n, err := strconv.Atoi(timeoutAsStr)
	if err != nil {
		log.Fatal(err)
	}
	// Create context with timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(n),
	)
	channel := make(chan any)
	// At the end of main, must close channel and cancel context
	defer close(channel)
	defer cancel()
	// Add two workers to the wait group
	wg.Add(2)
	// Start sending and receiving data
	go receiver(ctx, &wg, channel)
	go sender(ctx, &wg, channel)
	// Start waiting
	wg.Wait()
}
