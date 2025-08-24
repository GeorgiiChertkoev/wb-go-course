package main

import (
	"fmt"
	"sync"
)

func doubleAndGiveItToTheNext(ch chan int) chan int {
	doubled := make(chan int)
	go func() {
		defer close(doubled)
		for n := range ch {
			doubled <- n * 2
		}
	}()
	return doubled
}

func printIntegers(ch chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	baseChan := make(chan int)
	doubled := doubleAndGiveItToTheNext(baseChan)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); printIntegers(doubled) }()

	go func() {
		defer wg.Done()
		defer close(baseChan)
		for i := range 40 {
			baseChan <- i // отправляем числа для удвоения
		}
	}()

	wg.Wait()
}
