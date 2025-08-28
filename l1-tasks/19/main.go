package main

import "fmt"

func Reverse(s string) string {
	letters := []rune(s)

	for i := 0; i < len(letters)/2; i++ {
		letters[i], letters[len(letters)-i-1] = letters[len(letters)-i-1], letters[i]
	}
	return string(letters)
}

func main() {
	s := string("привет")
	fmt.Printf("was: %s\ngot: %s\n", s, Reverse(s))
}
