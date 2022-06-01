package main

import "fmt"

func binarySearch(list []int, item int) int {
	// Left and right borders
	left, right := 0, len(list)-1
	// Do while distance between borders >= 1
	for left <= right {
		midIdx := (left + right) / 2
		if list[midIdx] == item {
			return midIdx
			// If middle element greater then wanted - move right border
		} else if list[midIdx] > item {
			right = midIdx - 1
			// If middle element less then wanted - move left border
		} else {
			left = midIdx + 1
		}
	}
	// If no such element in the array return -1 as idx
	return -1
}

func main() {
	myList := []int{1, 3, 5, 6, 8, 17, 19, 21}

	fmt.Println(binarySearch(myList, 3))
	fmt.Println(binarySearch(myList, 17))
	fmt.Println(binarySearch(myList, 4))
}
