package main

import "fmt"

func Intersect[T comparable](a, b []T) []T {
	if len(a) > len(b) { // хотим добавить в мапу меньше ключей
		a, b = b, a
	}

	m := make(map[T]bool)
	for _, v := range a {
		m[v] = true
	}

	res := make([]T, 0, len(m))

	for _, v := range b {
		if m[v] {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 2, 4}

	fmt.Printf("%v\n", Intersect(a, b))
}
