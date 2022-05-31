package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// Worker function template. It sending date from channel to stdout and dies by
// context calling
func worker(ctx context.Context, readChannel <-chan any, number int) {
	// Endlessly listening readChannel and waiting ctx.Done()
	for {
		select {
		case value := <-readChannel:
			io.WriteString(
				os.Stdout,
				fmt.Sprintf("worker number %d got value: %v\n", number, value),
			)
			time.Sleep(time.Second * 1)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	// Parse second console arg for get number of workers
	if len(os.Args) < 2 {
		log.Fatal("Enter number of workers after program name.")
	}
	workersNumberAsStr := os.Args[1]
	workersNumberAsInt, err := strconv.Atoi(workersNumberAsStr)
	if err != nil {
		log.Fatal(err)
	}
	// Create context with cancel function
	ctx, cancel := context.WithCancel(context.Background())
	// Call cancel function in the end of main
	defer cancel()
	// Create channel for workers
	channel := make(chan any, workersNumberAsInt)
	// Create entered number of workers
	for i := 0; i < workersNumberAsInt; i++ {
		go worker(ctx, channel, i)
	}
	// Endlessly sending data to channel
	for {
		channel <- time.Now()
	}
}
