package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	runeSlice := make([]rune, 100)
	copy(runeSlice, []rune(v)[:100])
	justString = string(runeSlice)
}

func main() {
	someFunc()
}
