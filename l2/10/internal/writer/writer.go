package writer

import (
	"bufio"
	"fmt"
	"os"
)

// пишет выходные данные
type Writer struct {
	OutputFile  string
	OnlyUniques bool
}

func (w *Writer) WriteStrings(ch <-chan string) {
	var out *os.File
	var err error
	var prev string

	if w.OutputFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(w.OutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot create output file %s: %v", w.OutputFile, err)
		}
		defer out.Close()
	}

	bw := bufio.NewWriter(out)
	defer bw.Flush()

	for s := range ch {
		if w.OnlyUniques && prev == s {
			continue
		}
		bw.WriteString(s)
		bw.WriteByte('\n')
		prev = s
	}
}
