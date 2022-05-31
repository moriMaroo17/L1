package main

import "fmt"

type human struct {
	name string
	age  uint
}

func (h *human) getName() string {
	return h.name
}

func (h *human) getAge() uint {
	return h.age
}

type action struct {
	actionType string
	human
}

func main() {
	action := action{
		actionType: "move",
		human:      human{name: "alice", age: 21},
	}
	fmt.Printf("%s\n", action.getName())
	fmt.Printf("%d\n", action.getAge())
}
