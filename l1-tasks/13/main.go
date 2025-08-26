package main

import "fmt"

func main() {
	a, b := 1, 2

	// работает засчет обратимости XOR (a ^ b ^ a = b)
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)
}
