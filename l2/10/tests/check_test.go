package tests

import (
	"strings"
	"testing"

	"unix-sort/internal/args"
	"unix-sort/internal/sorter"
)

func TestCheckSortedColumn2Numeric(t *testing.T) {
	opts := args.ParseArgs(strings.Split(`-c -k 2 -n test_data/sorted_k2.csv`, " "))

	isSorted, _, err := sorter.Check(opts)
	if err != nil {
		t.Fatalf("expected file to be sorted, got error: %v", err)
	}
	if !isSorted {
		t.Errorf("given file is sorted, but check returned false")

	}
}
