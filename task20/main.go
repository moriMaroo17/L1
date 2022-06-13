package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(value string) {
	// Split by whitespaces
	splitedValue := strings.Fields(value)
	// Create builder
	builder := strings.Builder{}
	// Iterate from end to start of words array
	for i := len(splitedValue) - 1; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("%s ", splitedValue[i]))
	}
	fmt.Println(strings.TrimSpace(builder.String()))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Endlessly scan stdin
	for scanner.Scan() {
		reverse(scanner.Text())
	}
}
