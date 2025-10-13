package main

import (
	"fmt"
	"testing"
	"time"
)

func TestOrChannels(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-Or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	if time.Since(start) > 2*time.Second {
		t.Errorf("channel took too long, expected ~1sec got: %v sec", time.Since(start).Seconds())
	}
	fmt.Printf("done after %v", time.Since(start))
}
