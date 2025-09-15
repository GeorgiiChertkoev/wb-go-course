package main

import (
	"testing"
)

func TestSizingFunc(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"a3b", 4},
		{"a10", 10},
		{"abc", 3},
		{"a23qb3", 27},
		{"\ab", 2},
		{`\a\11`, 2},
		{`\a10\015`, 25},
		{`\122я2`, 24},
		{``, 0},
	}

	for _, tt := range tests {
		got := futureStringLen(tt.s)
		if got != tt.want {
			t.Errorf("input: %s, got %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestBrokenStrings(t *testing.T) {
	tests := []string{
		"1asd",
		"123",
		`4\12`,
	}
	for _, input := range tests {
		s, err := Unpack(input)
		if err == nil {
			t.Errorf("input: %s, expected err, got %s", input, s)
		}
	}
}

func TestSimpleUnpack(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
		isError  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"aq10b3", "aqqqqqqqqqqbbb", false},
	}

	for _, tt := range tests {
		got, err := Unpack(tt.input)
		if (err != nil) != tt.isError {
			t.Errorf("input: %s, got err: %v, expected error: %T", tt.input, err, tt.isError)
			continue
		}
		if got != tt.expected {
			t.Errorf("input: %s, got %v, expected %v", tt.input, got, tt.expected)
		}
	}
}

func TestEscapeSequences(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
		isError  bool
	}{
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`\1\2\3`, "123", false},
		{`\1\210\00`, "12222222222", false},
	}

	for _, tt := range tests {
		got, err := Unpack(tt.input)
		if (err != nil) != tt.isError {
			t.Errorf("input: %s, got err: %v, expected error: %T", tt.input, err, tt.isError)
			continue
		}
		if got != tt.expected {
			t.Errorf("input: %s, got %v, expected %v", tt.input, got, tt.expected)
		}
	}
}
