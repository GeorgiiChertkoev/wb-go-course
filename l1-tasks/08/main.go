package main

import "fmt"

const countingBitsFromOne = true

func changeBit(num int64, position int, value bool) int64 {
	if countingBitsFromOne {
		position-- // внутри функции я считаю с нуля
	}
	if value {
		num |= 1 << position
	} else {
		num &= -1 ^ (1 << position) // у -1 все биты единички в двоичном представлении
	}
	return num
}

func main() {
	n := int64(5)
	fmt.Printf("was: %x, got %x", n, changeBit(n, 1, false))
}
