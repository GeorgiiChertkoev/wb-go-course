package sorter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unix-sort/internal/args"
	"unix-sort/internal/comparator"
)

type SortDiagnose struct {
	LineNumber int
	Previous   string
	OutOfOrder string
}

func (d SortDiagnose) Print() {
	fmt.Printf("Out of order in line %d, \"%s\" should be before \"%s\"", d.LineNumber, d.OutOfOrder, d.Previous)
}

func Check(opts args.SortOptions) (isSorted bool, diagnose SortDiagnose, err error) {
	if len(opts.Files) > 1 {
		err = fmt.Errorf("can't take more than 1 file for checking (got %d: %v)", len(opts.Files), opts.Files)
		return
	}

	comparer, err := comparator.BuildComparator(opts)
	if err != nil {
		return
	}
	var r io.ReadCloser
	if len(opts.Files) == 0 {
		r = io.NopCloser(os.Stdin)
	} else {
		r, err = os.Open(opts.Files[0])
		if err != nil {
			return
		}
	}
	scanner := bufio.NewScanner(r)

	var cur string
	var prev string
	scanner.Scan()
	prev = scanner.Text()
	i := 0
	for scanner.Scan() {
		cur = scanner.Text()
		if comparer(cur, prev) { // cur < prev
			isSorted = false
			diagnose.LineNumber = i
			diagnose.Previous = prev
			diagnose.OutOfOrder = cur
			return
		}
		prev = cur
		i++
	}
	isSorted = true
	return
}
