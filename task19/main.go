package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(value string) {
	// Convert to rune array
	runeValue := []rune(value)
	result := make([]rune, len(runeValue))
	// Iterate from end to start of rune array
	for i := len(runeValue) - 1; i >= 0; i-- {
		result = append(result, runeValue[i])
	}
	fmt.Println(string(result))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Endlessly scan stdin
	for scanner.Scan() {
		reverse(scanner.Text())
	}
}
