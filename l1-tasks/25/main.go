package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepUsingFor(d time.Duration) {
	startTime := time.Now()

	for time.Since(startTime) < d {
	}
}

func sleepUsingChan(d time.Duration) {
	<-time.After(d)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sleepUsingFor(3 * time.Second)
		fmt.Println("awake after using for")
	}()

	go func() {
		defer wg.Done()
		sleepUsingChan(3 * time.Second)
		fmt.Println("awake after using chan")
	}()

	wg.Wait()
}
