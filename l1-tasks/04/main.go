package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const numberOfWorkers = 3

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup // c помощью wait group дождемся завершения всех горутин

	for i := range numberOfWorkers { // запускаем работников
		wg.Add(1)
		go func() {
			id := i
			defer wg.Done()
			for {
				select {
				/*
					благодаря такому подходу итерация цикла закончится и только после этого
					закончится горутина
				*/
				case <-ctx.Done():
					// ждет выполнится когда отменится контекст
					// (при его отмене канал закроется и придет нулевое значение)
					fmt.Printf("worker %d: done\n", id)
					return
				default:
					// имитируем какую-то работу
					time.Sleep(3 * time.Second)
					fmt.Printf("worker %d: did job\n", id)
				}
			}
		}()
	}

	wg.Wait() // дождемся завершения горутин
}
