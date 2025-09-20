package sorter

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// элемент кучи
type fileEntry struct {
	line  string // текущая строка
	index int    // индекс файла
}

type minHeap struct {
	data     []fileEntry
	comparer func(a, b string) bool
}

func (h *minHeap) Len() int           { return len(h.data) }
func (h *minHeap) Less(i, j int) bool { return h.comparer(h.data[i].line, h.data[j].line) }
func (h *minHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *minHeap) Push(x any)         { h.data = append(h.data, x.(fileEntry)) }
func (h *minHeap) Pop() any {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}

// mergeFiles открывает все временные файлы, делает k-way merge и отправляет результат в канал
func mergeFiles(files []string, out chan<- string) {
	defer close(out)

	// открыть все файлы
	scanners := make([]*bufio.Scanner, 0, len(files))
	h := &minHeap{}
	heap.Init(h)

	for i, fpath := range files {
		f, err := os.Open(fpath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v", fpath, err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f) // possible feature: add bigger buffer for optimization?
		scanners = append(scanners, scanner)
		if scanner.Scan() {
			heap.Push(h, fileEntry{
				line:  scanner.Text(),
				index: i,
			})
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "scan %s: %v", fpath, err)
		}
	}

	// k-way merge
	for h.Len() > 0 {
		// достаём минимальную строку
		entry := heap.Pop(h).(fileEntry)
		out <- entry.line

		// достаем следующую строку из того же файла
		if scanners[entry.index].Scan() {
			entry.line = scanners[entry.index].Text()
			heap.Push(h, entry)
		} else {
			if err := scanners[entry.index].Err(); err != nil {
				fmt.Fprintf(os.Stderr, "scan error: %v", err)
			}
		}
	}
}
