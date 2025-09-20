package sorter

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"sync"
	"unix-sort/internal/args"
	"unix-sort/internal/comparator"
	"unix-sort/internal/reader"
	"unix-sort/internal/writer"
)

// больше одной горутины на файл = бессмысленно из-за syscall-ов
// поэтому одна горутина для чтения файла и одна для записи
// GOMAXPROCS горутин для сортировки
// будем общаться через chan []string где каждый слайс не больше buffer-size
// сортим слайс и кладем его в временный файл
// когда мы всё прочитали, канал закрывается, мы досортируем все файлы и сохраним их в времнную директорию
// дальше сделаем k-way merge. будем считывать строки из каждого файла в встроенный буффер сканера
// дальше добавим в мин кучу первый элемент из каждого файла
// и будем доставать из нее значения, заменяя их следующими из соотв очереди
func Sort(opts args.SortOptions) error {
	comparer, err := comparator.BuildComparator(opts)
	if err != nil {
		return err
	}
	if opts.CheckSorted {
		isSorted, diagnose, err := Check(opts)
		if err != nil {
			return err
		}
		if !isSorted {
			diagnose.Print()
		}
		return nil
	}

	slices := make(chan []string, runtime.GOMAXPROCS(0))
	sortedSlices := make(chan []string, runtime.GOMAXPROCS(0))

	go reader.Read(opts, slices)

	var wgWorkers sync.WaitGroup

	for range runtime.GOMAXPROCS(0) {
		wgWorkers.Add(1)
		go func() {
			defer wgWorkers.Done()
			worker(slices, sortedSlices, comparer)
		}()
	}
	go func() {
		// закроем канал когда все воркеры закончат
		wgWorkers.Wait()
		close(sortedSlices)
	}()

	tempFiles := writeTempFiles(sortedSlices, opts.TempDir)
	if !opts.KeepTemps {
		defer cleanUp(tempFiles)
	}
	out := make(chan string, 512)
	go mergeFiles(tempFiles, out, comparer)
	w := writer.Writer{
		OutputFile:  opts.OutputFile,
		OnlyUniques: opts.Unique,
	}
	w.WriteStrings(out)

	return nil
}

func cleanUp(files []string) {
	for _, f := range files {
		os.Remove(f)
	}
}

func worker(slices <-chan []string, sortedSlices chan<- []string, comparer func(a, b string) bool) {
	for slice := range slices {
		sort.Slice(slice, func(i, j int) bool {
			return comparer(slice[i], slice[j])
		})
		sortedSlices <- slice
	}
}

func writeTempFiles(slices <-chan []string, dir string) []string {
	var files []string
	i := 0

	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "create dir %s: %v", dir, err)
	}

	for chunk := range slices {
		f, err := os.CreateTemp(dir, "chunk_*.tmp")
		name := f.Name()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create file %s: %v", name, err)
		}
		w := bufio.NewWriter(f)
		for _, line := range chunk {
			w.WriteString(line)
			w.WriteByte('\n')
		}
		w.Flush()
		f.Close()
		files = append(files, name)
		i++
	}

	return files
}
