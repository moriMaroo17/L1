package main

import (
	"fmt"
	"log"
)

func setToOne(num int64, pos uint) int64 {
	return num | (1 << pos)
}

func setToZero(num int64, pos uint) int64 {
	// set bit to zero if right part bit is one. Else, take value from left
	return num &^ (1 << pos)
}

func setBit(num int64, pos uint, val uint) int64 {
	if val == 0 {
		return setToZero(num, pos)
	} else if val == 1 {
		return setToOne(num, pos)
	} else {
		log.Fatal("bit can be set only to one or zero")
		return 0
	}
}

func main() {
	var number int64 = 132
	fmt.Printf("Origin: %b\n", number)
	fmt.Printf("Set 4th bit to 1: %b\n", setBit(number, 4, 1))
	fmt.Printf("Set 2nd bit to 0: %b\n", setBit(number, 2, 0))
}
