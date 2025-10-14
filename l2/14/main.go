package main

import (
	"fmt"
	"time"
)

func Or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(<-chan interface{})
	for _, ch := range channels {
		res = orTwoChannels(res, ch)
	}

	return res
}

func OrBySelect(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})

	go func() {
		for {
			for _, c := range channels {
				select {
				case v, ok := <-c:
					if !ok {
						close(res)
						return
					}
					res <- v
				default:
				}
			}
		}
	}()

	return res
}

func orTwoChannels(ch1, ch2 <-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					close(res)
					return
				}
				res <- v
			case v, ok := <-ch2:
				if !ok {
					close(res)
					return
				}
				res <- v
			}
		}
	}()
	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-orTwoChannels(
		sig(3*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}
