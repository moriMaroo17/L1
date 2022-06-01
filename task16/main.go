package main

import "fmt"

func quickSort(arr []int) []int {
	// Base case
	if len(arr) < 2 {
		return arr
	}
	// Recursive case
	var less, greater []int
	// Кeference element selection
	pivot := arr[0]
	for _, item := range arr[1:] {
		if item < pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	fmt.Println(quickSort([]int{3, -3, 9, 33, 11, 4, 5, -5, 0, -5, 3, 6}))
}
