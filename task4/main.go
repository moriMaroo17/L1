package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// Worker function template. It sending date from channel to stdout and dies
// with closing channel
func worker(readChannel <-chan any, number int) {
	// Notify wait group about done work, and send message about death to stdout
	defer io.WriteString(os.Stdout, fmt.Sprintf("worker number %d died\n", number))
	// Endlessly listening readChannel while it's open
	for value := range readChannel {
		io.WriteString(
			os.Stdout,
			fmt.Sprintf("worker number %d got value: %v\n", number, value),
		)
		time.Sleep(time.Second)
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
	// Create cancel channel
	cancelChannel := make(chan os.Signal, 1)
	// Notify cancel channel about keyboard Interrupt
	signal.Notify(cancelChannel, os.Interrupt)
	// Call cancel function in the end of main
	defer func() {
		signal.Stop(cancelChannel)
	}()

	// Create channel for workers
	channel := make(chan any, workersNumberAsInt)
	// Create entered number of workers
	for i := 0; i < workersNumberAsInt; i++ {
		go worker(channel, i)
	}
	// Endlessly sending data to channel and wating interrupt
	for {
		select {
		case <-cancelChannel:
			close(channel)
			return
		default:
			channel <- time.Now().UnixNano()
		}
	}
}
