package main

import (
	"fmt"
	"sync"
)

func main() {
	numberOfWorkers := 50
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)
	ch := make(chan any, numberOfWorkers) // канал с буффером на кол-во горутин

	for range numberOfWorkers {
		go func() {
			defer wg.Done()
			for {
				value, ok := <-ch
				if !ok { // выйдем из горутины когда значения в канале закончатся и он будет закрыт
					break
				}
				fmt.Println(value)
			}
		}()
	}

	for i := range 100 {
		ch <- i // отправляем данные в канал
	}
	close(ch)

	wg.Wait() // ждем когда все горутины закончат
}
