package main

import "fmt"

func distributer(arr []float64) map[int][]float64 {
	result := make(map[int][]float64)
	// for each elem in the arr call nearest() and append with such key
	// to the result map
	for _, elem := range arr {
		near := nearest(elem)
		result[near] = append(result[near], elem)
	}
	return result
}

// Function, which finds neared border
func nearest(value float64) int {
	// Takes the first numbers digit
	firstDigit := int(value) / 10
	var border1, border2 int
	// Way to find borders depends on the sign of the number
	if value > 0 {
		border1, border2 = firstDigit*10, (firstDigit+1)*10
		if float64(border2)-value > value-float64(border1) {
			return border1
		}
		return border2
	}
	border1, border2 = (firstDigit-1)*10, firstDigit*10
	if float64(border1)-value > value-float64(border2) {
		return border1
	}
	return border2
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := distributer(arr)
	fmt.Printf("%v\n", result)
}
