package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 123131231212312
	b := 2311231231231

	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))

	fmt.Println(new(big.Int).Add(bigA, bigB))
	fmt.Println(new(big.Int).Sub(bigA, bigB))
	fmt.Println(new(big.Int).Mul(bigA, bigB))
	fmt.Println(new(big.Int).Div(bigA, bigB))
}
