package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unix-sort/internal/args"
)

// считывание текста из файлов и stdin

func Read(opts args.SortOptions, ch chan<- []string) {
	defer close(ch)

	var chunk []string
	var size int

	flush := func() {
		if len(chunk) > 0 {
			tmp := make([]string, len(chunk))
			copy(tmp, chunk)
			ch <- tmp
			chunk = chunk[:0]
			size = 0
		}
	}

	// функция добавления строки в общий буфер
	addLine := func(line string) {
		lineSize := len(line) + 1
		if int64(size+lineSize) > opts.BufSize && len(chunk) > 0 {
			flush()
		}
		chunk = append(chunk, line)
		size += lineSize
	}

	// читаем stdin или файлы
	if len(opts.Files) == 0 {
		if err := readFrom(io.NopCloser(os.Stdin), addLine); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		}
		return
	}

	for _, p := range opts.Files {
		f, err := os.Open(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", p, err)
			continue
		}
		if err := readFrom(f, addLine); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", p, err)
		}
		_ = f.Close()
	}

	flush()
}

func readFrom(r io.ReadCloser, emit func(string)) error {
	defer r.Close()

	scanner := bufio.NewScanner(r)

	const maxLine = 16 * 1024 * 1024 // 16 MB
	buf := make([]byte, 0, 64*1024)  // изначально буффер на 64кб
	scanner.Buffer(buf, maxLine)

	for scanner.Scan() {
		emit(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
