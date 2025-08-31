package main

import (
	"fmt"
	"math/big"
)

func DoMath(a, b *big.Int) {
	var z big.Int
	fmt.Printf("a+b: %s\n", z.Add(a, b))
	fmt.Printf("a-b: %s\n", z.Sub(a, b))
	fmt.Printf("a*b: %s\n", z.Mul(a, b))
	fmt.Printf("a/b: %s\n", z.Div(a, b))
}

func main() {
	DoMath(big.NewInt(1<<50), big.NewInt(1<<40))
}
