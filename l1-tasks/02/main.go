package main

import (
	"fmt"
	"sync"
)

func solution1(nums []int) {
	var wg sync.WaitGroup
	for _, value := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(value)
	}
	wg.Wait()
}

func solution2(nums []int) {
	results := make(chan int)
	for _, value := range nums {
		go func(num int) {
			results <- num * num
		}(value)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-results)
	}

}

func main() {
	array := []int{2, 4, 6, 8, 10}
	solution1(array)
	fmt.Println()
	solution2(array)

}
