package comparator

import (
	"testing"
	"unix-sort/internal/args"
)

type builderTest struct {
	Opts    args.SortOptions
	A, B    string
	CompRes bool
}

func TestComparatorBuilding(t *testing.T) {
	tests := []builderTest{
		{
			Opts:    args.SortOptions{},
			A:       "abc",
			B:       "bbc",
			CompRes: true,
		},
		{
			Opts: args.SortOptions{
				Reverse: true,
			},
			A:       "abc",
			B:       "bbc",
			CompRes: false,
		},
		{
			Opts: args.SortOptions{
				Numeric: true,
				Reverse: true,
			},
			A:       "98",
			B:       "0099",
			CompRes: false,
		},
		{
			Opts: args.SortOptions{
				HumanNumeric: true,
				KeyColumn:    2,
				Separators:   "\t",
			},
			A:       "1\t98M",
			B:       "2\t0099K",
			CompRes: false,
		},
		{
			Opts: args.SortOptions{
				Month:      true,
				Reverse:    true,
				KeyColumn:  3,
				Separators: " ",
			},
			A:       "a 2 JUN",
			B:       "3 b DEC",
			CompRes: false,
		},
		{
			Opts: args.SortOptions{
				Numeric:      true,
				Reverse:      true,
				IgnoreBlanks: true,
			},
			A:       "     12",
			B:       " 13 ",
			CompRes: false,
		},
	}

	for _, tt := range tests {
		comparer, err := BuildComparator(tt.Opts)
		if err != nil {
			t.Errorf("error making comparator: %v with opts: %v", err, tt.Opts)
		}
		res := comparer(tt.A, tt.B)
		if res != tt.CompRes {
			t.Errorf("compared %s and %s with params: %#v; Expected %t, got %t", tt.A, tt.B, tt.Opts, tt.CompRes, res)
		}
	}
}
