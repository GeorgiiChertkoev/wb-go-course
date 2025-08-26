package main

import "fmt"

func PrintType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("got int\n")
	case string:
		fmt.Printf("got string\n")
	case bool:
		fmt.Printf("got bool\n")
	case chan interface{}:
		fmt.Printf("got chan interface\n")
	default:
		fmt.Printf("Was not ready for type: %T", v)
	}
}

func main() {
	var a any
	a = false
	PrintType(a)
	a = 1
	PrintType(a)
	a = string("qwerty")
	PrintType(a)
	a = make(chan any)
	PrintType(a)
}
