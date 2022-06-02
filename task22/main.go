package main

import (
	"errors"
	"fmt"
	"math/big"
)

func multiple(num1, num2 *big.Int) *big.Int {
	return new(big.Int).Mul(num1, num2)
}

func add(num1, num2 *big.Int) *big.Int {
	return new(big.Int).Add(num1, num2)
}

func divide(num1, num2 *big.Int) (*big.Int, error) {
	if num2 == big.NewInt(0) {
		return nil, errors.New("divide by zero")
	}
	return new(big.Int).Div(num1, num2), nil
}

func subtraction(num1, num2 *big.Int) *big.Int {
	return new(big.Int).Sub(num1, num2)
}

func main() {
	bigInt1 := big.NewInt(int64(1 << 50))
	bigInt2 := big.NewInt(int64(1 << 40))

	fmt.Println(multiple(bigInt1, bigInt2))
	fmt.Println(add(bigInt1, bigInt2))
	fmt.Println(divide(bigInt1, bigInt2))
	fmt.Println(subtraction(bigInt1, bigInt2))
}
