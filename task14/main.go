package main

import "fmt"

func getType(variable interface{}) {
	// Type switch
	switch variable.(type) {
	case int:
		fmt.Print("This is int\n")
	case string:
		fmt.Print("This is string\n")
	case bool:
		fmt.Print("This is bool\n")
	case chan any:
		fmt.Print("This is channel\n")
	}
}

func main() {
	// Check types
	getType(1)
	getType("foo")
	getType(false)
	getType(make(chan interface{}))
}
