package main

import (
	"fmt"
	"sort"
)

func roundToTen(n float64) int {
	return int(n) / 10 * 10
}

func main() {
	tempretures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float64)

	for _, t := range tempretures {
		m[roundToTen(t)] = append(m[roundToTen(t)], t)
	}

	// сортирую ключи чтобы вывод был читабельнее
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, m[k])
	}
	// вывод:
	// -20: [-25.4 -27 -21]
	// 10: [13 19 15.5]
	// 20: [24.5]
	// 30: [32.5]
}
