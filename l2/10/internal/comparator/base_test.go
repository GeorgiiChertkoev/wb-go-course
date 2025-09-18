package comparator

import "testing"

type Test struct {
	A      string
	B      string
	IsLess bool
}

func TestLexicographic(t *testing.T) {
	tests := []Test{
		{"a", "b", true},
		{"bb", "ba", false},
		{"bb", "bb", false},
		{"a1b", "a2", true},
		{"A", "a", true},
	}

	for _, tt := range tests {
		compareRes := lexicographic(tt.A, tt.B)
		if compareRes != tt.IsLess {
			t.Errorf("lexi compared %s < %s, got %t, expected %t", tt.A, tt.B, compareRes, tt.IsLess)
		}
	}
}

func TestNumeric(t *testing.T) {
	tests := []Test{
		{"1", "2", true},
		{"123", "999", true},
		{"0.1", "0.01", false},
		{"", "1", true}, // An empty number is treated as ‘0’.
		{"12", "qwerty", false},
		{"-12", "12", true},
		{"000011", "11", false},
		{"000011", "11", false},
		{"000", "0", false},
		{"-0", "0", false},
		{"1_0", "9", false},
	}

	for _, tt := range tests {
		compareRes := numeric(tt.A, tt.B)
		if compareRes != tt.IsLess {
			t.Errorf("numeric compared %s < %s, got %t, expected %t", tt.A, tt.B, compareRes, tt.IsLess)
		}
	}
}

func TestMonth(t *testing.T) {
	tests := []Test{
		{"a", "JAN", true},
		{"JAN", "FEB", true},
		{"JAN", "   FEB", true},
		{"DEC", "OCT", false},
		{"SEP", "NOV", true},
	}

	for _, tt := range tests {
		compareRes := month(tt.A, tt.B)
		if compareRes != tt.IsLess {
			t.Errorf("month compared %s < %s, got %t, expected %t", tt.A, tt.B, compareRes, tt.IsLess)
		}
	}
}

func TestHumanNumeric(t *testing.T) {
	tests := []Test{
		// ведет себя как numeric если нет + в начале и буквы в конце
		{"1", "2", true},
		{"123", "999", true},
		{"0.1", "0.01", false},
		{"", "1", true}, // An empty number is treated as ‘0’.
		{"12", "qwerty", false},
		{"-12", "12", true},
		{"000011", "11", false},
		{"000011", "11", false},
		{"000", "0", false},
		{"1_0", "9", false},
		// только для humanNumeric
		{"10K", "1M", true},
		{"10K", "-1M", false},
		{"1191991", "1K", true},
		{"2M", "+1G", true},
		{"2Q", "1Y", false},
		{"20Q", "19Q", false},
		{"101T", "102T", true},
	}

	for _, tt := range tests {
		compareRes := humanNumeric(tt.A, tt.B)
		if compareRes != tt.IsLess {
			t.Errorf("human numeric compared %s < %s, got %t, expected %t", tt.A, tt.B, compareRes, tt.IsLess)
		}
	}
}
