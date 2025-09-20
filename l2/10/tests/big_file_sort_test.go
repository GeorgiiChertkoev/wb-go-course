package tests

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"unix-sort/internal/args"
	"unix-sort/internal/sorter"
)

func isSorted(data []string) bool {
	return sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func TestSort_SplitsIntoMultipleChunks(t *testing.T) {
	// создаём большой временный файл с 10000 строк
	tmpFile := t.TempDir() + "/big_input.txt"
	f, err := os.Create(tmpFile)
	if err != nil {
		t.Fatal(err)
	}
	for i := 10000; i >= 1; i-- { // в обратном порядке
		fmt.Fprintf(f, "line-%05d\n", i)
	}
	f.Close()

	tmpOut := t.TempDir() + "/out.txt"
	// специально очень маленький буффер чтобы симулировать работу как для больших файлов
	opts := args.ParseArgs(strings.Split("-S 1024 "+tmpFile, " "))
	opts.OutputFile = tmpOut

	if err := sorter.Sort(opts); err != nil {
		t.Fatalf("sort failed: %v", err)
	}

	data, err := os.ReadFile(tmpOut)
	if err != nil {
		t.Fatalf("cannot read output: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	if len(lines) != 10000 {
		t.Errorf("expected 10000 lines, got %d", len(lines))
	}

	if !isSorted(lines) {
		t.Errorf("output is not sorted")
	}
}
