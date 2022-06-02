package main

import "fmt"

func deleteElement(arr []int, number int) []int {
	result := make([]int, len(arr)-1)
	copy(result, append(arr[:number], arr[number+1:]...))
	return result
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(cap(arr))
	fmt.Println(arr)
	arr = deleteElement(arr, 5)
	fmt.Println(cap(arr))
	fmt.Println(arr)
}
