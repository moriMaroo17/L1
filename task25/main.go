package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	// Starts timer
	timer := time.NewTimer(duration)
	// Waiting for timer stop
	<-timer.C
}

func main() {
	fmt.Println(time.Now().Second())
	// Test native Sleep function
	time.Sleep(time.Second)
	fmt.Println(time.Now().Second())
	// Test mySleep function
	mySleep(time.Second)
	fmt.Println(time.Now().Second())
}
