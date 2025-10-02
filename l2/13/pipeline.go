package main

import (
	"fmt"
	"io"
	"strings"
)

func Filter(ch <-chan string, substr string) <-chan string {
	filtered := make(chan string, pipelineChanSize)
	go func() {
		defer close(filtered)
		for s := range ch {
			if strings.Contains(s, substr) {
				filtered <- s
			}
		}
	}()

	return filtered
}

func SplitIntoFields(ch <-chan string, delimeter string) <-chan []string {
	splited := make(chan []string, pipelineChanSize)
	go func() {
		defer close(splited)
		for s := range ch {
			splited <- strings.Split(s, delimeter)
		}
	}()
	return splited
}

func FilterFields(ch <-chan []string, fieldsId []int) <-chan []string {
	if len(fieldsId) == 0 {
		return ch
	}
	filtered := make(chan []string, pipelineChanSize)
	go func() {
		defer close(filtered)
		for fields := range ch {
			res := make([]string, 0)
			for _, id := range fieldsId {
				if id-1 >= len(fields) {
					continue
				}
				res = append(res, fields[id-1])
			}
			filtered <- res
		}
	}()
	return filtered

}

// blocking
func PrintFields(ch <-chan []string, delimeter string, writer io.WriteCloser) {
	defer writer.Close()
	for lines := range ch {
		fmt.Fprintln(writer, strings.Join(lines, delimeter))
	}
}
