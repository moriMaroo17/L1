package main

import (
	"fmt"
	"strings"
)

func findDubles(value string) bool {
	runeValue := []rune(strings.ToLower(value))
	chars := make(map[rune]struct{})
	for _, char := range runeValue {
		if _, ok := chars[char]; ok {
			return !ok
		}
		chars[char] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(findDubles("abcd"))
	fmt.Println(findDubles("abCdefAaf"))
	fmt.Println(findDubles("aabcd"))
}
