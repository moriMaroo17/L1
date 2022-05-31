package main

import "fmt"

// Empty struct for map 'data'
type void struct{}

// Set implementation
type set struct {
	data map[string]void
}

// Method for put value into set
func (s *set) put(value string) {
	s.data[value] = void{}
}

// Method for check containing value into the set
func (s *set) checkContain(value string) bool {
	_, ok := s.data[value]
	return ok
}

func main() {
	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	stringSet := set{data: make(map[string]void)}
	for _, elem := range arr {
		stringSet.put(elem)
	}
	fmt.Printf("%v\n", stringSet)
}
