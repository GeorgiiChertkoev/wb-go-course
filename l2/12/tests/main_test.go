package main

import (
	"bytes"
	"go-grep/internal/grep"
	"go-grep/internal/options"
	"strings"
	"testing"
)

func runGrep(t *testing.T, input_args []string) string {
	t.Helper()

	// эмулируем командную строку
	args, err := options.ParseArgs(input_args)
	if err != nil {
		t.Fatalf("ParseArgs failed: %v", err)
	}

	results, err := grep.Grep(*args)
	if err != nil {
		t.Fatalf("Grep failed: %v", err)
	}

	var buf bytes.Buffer
	for _, res := range results {
		res.Print(&buf)
	}
	return buf.String()
}
func TestSimpleMatch(t *testing.T) {
	output := runGrep(t, []string{"John", "us_presidents.txt"})
	if !strings.Contains(output, "John Adams") {
		t.Errorf("expected match for 'John Adams', got:\n%s", output)
	}
}

func TestCaseInsensitive(t *testing.T) {
	output := runGrep(t, []string{"-i", "john", "us_presidents.txt"})
	if !strings.Contains(output, "John Quincy Adams") {
		t.Errorf("expected case-insensitive match, got:\n%s", output)
	}
}

func TestInvertMatch(t *testing.T) {
	output := runGrep(t, []string{"-v", "Trump", "us_presidents.txt"})
	if strings.Contains(output, "Donald Trump") {
		t.Errorf("expected inverted match to exclude Trump, got:\n%s", output)
	}
}

func TestCountOnly(t *testing.T) {
	output := strings.TrimSpace(runGrep(t, []string{"-c", "John", "us_presidents.txt"}))
	if output != "4" {
		t.Errorf("expected count=4, got: %q", output)
	}
}

func TestLineNumbers(t *testing.T) {
	output := runGrep(t, []string{"-n", "John", "us_presidents.txt"})
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if !strings.HasPrefix(lines[0], "2:") {
		t.Errorf("expected line number prefix '2:' for John Adams, got: %q", lines[0])
	}
}

func TestContextAfter(t *testing.T) {
	output := runGrep(t, []string{"-A", "1", "John", "us_presidents.txt"})
	if !strings.Contains(output, "Thomas Jefferson") {
		t.Errorf("expected context line after John Adams, got:\n%s", output)
	}
}

func TestContextBefore(t *testing.T) {
	output := runGrep(t, []string{"-B", "1", "John", "us_presidents.txt"})
	if !strings.Contains(output, "George Washington") {
		t.Errorf("expected context line before John Adams, got:\n%s", output)
	}
}

func TestFixedString(t *testing.T) {
	output := strings.TrimSpace(runGrep(t, []string{"-F", "John Adams", "us_presidents.txt"}))
	if output != "John Adams, 1797-1801" {
		t.Errorf("expected exact fixed string match, got:\n%s", output)
	}
}

func TestRegexYears(t *testing.T) {
	output := runGrep(t, []string{"18[0-9]{2}", "us_presidents.txt"})
	if !strings.Contains(output, "Thomas Jefferson") {
		t.Errorf("expected a president from 1800s, got:\n%s", output)
	}
	if strings.Contains(output, "George Washington") {
		t.Errorf("should not match 1700s, got:\n%s", output)
	}
}

func TestExactYearRange(t *testing.T) {
	output := runGrep(t, []string{"1945-", "us_presidents.txt"})
	if !strings.Contains(output, "Harry S. Truman") {
		t.Errorf("expected Truman 1945, got:\n%s", output)
	}
}

func TestOpenEndedTerm(t *testing.T) {
	output := runGrep(t, []string{"[0-9]{4}-$", "us_presidents.txt"})
	if !strings.Contains(output, "Donald Trump, 2025-") {
		t.Errorf("expected open-ended Donald Trump 2025-, got:\n%s", output)
	}
}

func TestMultipleDigitsRegex(t *testing.T) {
	output := runGrep(t, []string{"^[A-Za-z .]+, [0-9]{4}-[0-9]{4}$", "us_presidents.txt"})
	if !strings.Contains(output, "John Adams, 1797-1801") {
		t.Errorf("expected strict YYYY-YYYY match for John Adams, got:\n%s", output)
	}
	if strings.Contains(output, "Donald Trump, 2025-") {
		t.Errorf("should not match open-ended term, got:\n%s", output)
	}
}

func TestContextAroundRegex(t *testing.T) {
	output := runGrep(t, []string{"-C", "1", "196[0-9]", "us_presidents.txt"})
	if !strings.Contains(output, "Dwight David Eisenhower") {
		t.Errorf("expected context before Kennedy 1961, got:\n%s", output)
	}
	if !strings.Contains(output, "Lyndon Baines Johnson") {
		t.Errorf("expected context after Kennedy 1961, got:\n%s", output)
	}
}
