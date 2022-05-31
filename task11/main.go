package main

import "fmt"

// Empty struct for map 'data'
type void struct{}

// Set implementation
type set struct {
	data map[int]void
}

// Method for put value into set
func (s *set) put(value int) {
	s.data[value] = void{}
}

// Method for check containing value into the set
func (s *set) checkContain(value int) bool {
	_, ok := s.data[value]
	return ok
}

// Function for find intersection
func findIntersection(set1, set2 set) set {
	intersection := set{data: make(map[int]void)}
	for key := range set1.data {
		if set2.checkContain(key) {
			intersection.put(key)
		}
	}
	return intersection
}

func main() {
	set1 := set{data: make(map[int]void)}
	set2 := set{data: make(map[int]void)}
	// Put data into set1
	set1.put(31)
	set1.put(8)
	set1.put(19)
	set1.put(45)
	set1.put(32)
	set1.put(27)
	// Put data into set2
	set2.put(1)
	set2.put(6)
	set2.put(27)
	set2.put(4)
	set2.put(31)
	set2.put(8)

	fmt.Printf("%v\n", findIntersection(set1, set2))
}
