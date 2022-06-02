package main

import (
	"fmt"
	"strings"
)

func findDubles(value string) bool {
	value = strings.ToLower(value)
	for idx, char := range value {
		if result := strings.Contains(value[idx+1:], string(char)); result {
			return !result
		}
	}
	return true
}

func main() {
	fmt.Println(findDubles("abcd"))
	fmt.Println(findDubles("abCdefAaf"))
	fmt.Println(findDubles("aabcd"))
}
