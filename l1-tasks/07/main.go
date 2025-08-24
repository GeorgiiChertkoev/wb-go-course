package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	letters := string("ababcbaacbabcbaacbbacbaqbabbababbaba")
	m := make(map[rune]int)
	var mutex sync.Mutex
	ch := make(chan rune, runtime.NumCPU())
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, letter := range letters {
			ch <- rune(letter)
		}
	}()

	for _ = range runtime.NumCPU() { // запускаем столько же воркеров сколько ядер
		wg.Add(1)
		go func() {
			defer wg.Done()
			for letter := range ch {
				// с помощью мьютекса захватываем мапу так что только мы ей пользуемся в этот момент
				mutex.Lock()
				m[letter]++ // если нет ключа создастся ячейка с нулевым значением и мы сразу к ней прибавим
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()

	for k, v := range m {
		fmt.Printf("got %c %d times\n", k, v)
	}
}
