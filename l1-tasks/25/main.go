package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool)
	for _, v := range words {
		m[v] = true
	}

	uniqueWords := make([]string, 0, len(m))
	for k, _ := range m {
		uniqueWords = append(uniqueWords, k)
	}

	fmt.Printf("Uniques: %v\n", uniqueWords)
}
