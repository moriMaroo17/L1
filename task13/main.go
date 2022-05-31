package main

import "fmt"

func main() {
	var1, var2 := 15, 25
	fmt.Printf("var1: %d, var2: %d\n", var1, var2)
	var1, var2 = var2, var1
	fmt.Printf("var1: %d, var2: %d\n", var1, var2)
}
