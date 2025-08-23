package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func runGoroutineWithWG(wg *sync.WaitGroup, f func()) {
	// Навесит waitGroup на функцию не меняя ее сигнатуры
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func doWork() {
	// имитация работы
	time.Sleep(500 * time.Millisecond)
}

func stopByExternalFlag(done *bool) {
	for {
		if *done {
			fmt.Println("stopped by external flag")
			return
		} else {
			doWork()
		}
	}
}

func stopByChannel(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("stopped by channel")
			return
		default:
			doWork()
		}
	}
}

func stopByContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped by context")
			return
		default:
			doWork()
		}
	}
}

func stopByGoexit(ctx context.Context) {
	// помощью контекста буду контроллировать когда вызвать goexit
	defer fmt.Println("stopped by Goexit")
	for {
		select {
		case <-ctx.Done():
			runtime.Goexit()
		default:
			doWork()
		}
	}
	fmt.Println("This will never happen")
}

func stopByPanic(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Stopped panic: ", r)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			panic("paniced to exit")
		default:
			doWork()
		}
	}
	fmt.Println("This will never happen")
}

func stopByTimeout(lifetime time.Duration) {
	ch := time.After(lifetime)
	for {
		select {
		case <-ch:
			fmt.Println("stopped by timeout")
			return
		default:
			doWork()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	var done bool = false // для выхода по условию
	// не atomic так как только одна горутина изменяет значение, остальные(1) читают

	ctx, cancel := context.WithCancel(context.Background())
	doneCh := make(chan struct{})

	runGoroutineWithWG(&wg, func() { stopByExternalFlag(&done) })
	runGoroutineWithWG(&wg, func() { stopByChannel(doneCh) })
	runGoroutineWithWG(&wg, func() { stopByContext(ctx) })
	runGoroutineWithWG(&wg, func() { stopByGoexit(ctx) })
	runGoroutineWithWG(&wg, func() { stopByPanic(ctx) })

	runGoroutineWithWG(&wg, func() { stopByTimeout(30 * time.Second) }) // прекратится независимо от внешних условий

	// вызовем все способы остановки горутин(кроме таймаута) на ctrl+c
	shutdownCtx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	runGoroutineWithWG(&wg, func() {
		<-shutdownCtx.Done() // ждем прерывания
		done = true          // завершаем горутину по условию
		cancel()             // завершаем горутины по контексту/с контролем извне
		close(doneCh)        // завершаем горутину по каналу
	})

	fmt.Println("Started all goroutines")
	wg.Wait()
}
