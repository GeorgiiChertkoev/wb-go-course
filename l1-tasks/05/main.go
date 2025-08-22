package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const N = 5 // время работы программы в секундах

func sendToChannel(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- 42
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func readFromChannel(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Read from channel: %d\n", num)
		}
	}
}

func main() {
	// через N секунд контекст отменится и в select выберется выход из функции
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	ch := make(chan int) // канал для обмена без буффера
	wg.Add(2)
	go sendToChannel(ctx, ch, &wg)
	go readFromChannel(ctx, ch, &wg)

	wg.Wait()
}
