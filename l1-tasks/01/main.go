package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi %s", h.name)
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{
			name: "George",
		},
	}

	a.SayHi()
}
