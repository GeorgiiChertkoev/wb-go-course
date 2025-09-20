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

// основная логика сортировки
func Sort(opts args.SortOptions) error {
	/*

		будем общаться через chan []string где каждый слайс чуть больше 256мб
		("чуть больше" потому что будем считывать пока не найдем перенос строки и размер слайса выйдет за 1 гб)
		// далее 256мб можно поменять на buffer-size из параметров

		сортим слайс и кладем его в временный файл

		больше одной горутины на файл = бессмысленно из-за syscall-ов
		одна горутина для чтения файла
		GOMAXPROCS или -parallel  горутин для сортировки и записи

		когда мы всё прочитали, канал закрывается, мы досортируем все файлы и сохраним их в времнную директорию

		дальше сделаем k-way merge. будем считывать некоторое число строк из каждого файла в свою очередь
		дальше добавим в мин кучу первый элемент каждой очереди
		и будем доставать из нее значения, заменяя их следующими из соотв очереди

	*/

	comparer, err := comparator.BuildComparator(opts)
	if err != nil {
		return err
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
	out := make(chan string, 512)
	go mergeFiles(tempFiles, out)
	w := writer.Writer{OutputFile: opts.OutputFile}
	w.WriteStrings(out)

	fmt.Println(tempFiles)
	return nil
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
