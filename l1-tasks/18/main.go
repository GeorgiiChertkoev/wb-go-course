package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ConcurrentCounter struct {
	Counter atomic.Int64
}

func (c *ConcurrentCounter) Intcrement() {
	c.Counter.Add(1)
}

var c ConcurrentCounter

func main() {
	defer fmt.Printf("%v\n", c.Counter.Load())

	var wg sync.WaitGroup
	wg.Add(10)
	for _ = range 10 {
		go func() {
			defer wg.Done()
			for _ = range 10 {
				c.Intcrement()
			}
		}()
	}

	wg.Wait()

}
