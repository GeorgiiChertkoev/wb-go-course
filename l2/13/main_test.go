package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

type nopWriteCloser struct {
	*bytes.Buffer
}

func (nwc nopWriteCloser) Close() error { return nil }

func TestCut_SimpleFields(t *testing.T) {
	var out bytes.Buffer
	opts := CutOpts{
		InputFiles: []string{"test_data/data.csv"},
		Fields:     []int{1, 3},
		Delimeter:  "\t",
		Output:     nopWriteCloser{&out},
	}
	if err := Cut(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "a\tc\n1\t3\nx\tz\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestCut_RangeFields(t *testing.T) {
	var out bytes.Buffer
	opts := CutOpts{
		InputFiles: []string{"test_data/data.csv"},
		Fields:     []int{2, 3, 4},
		Delimeter:  "\t",
		Output:     nopWriteCloser{&out},
	}
	if err := Cut(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "b\tc\td\n2\t3\t4\ny\tz\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestCut_CombinationFields(t *testing.T) {
	var out bytes.Buffer
	opts := CutOpts{
		InputFiles: []string{"test_data/data.csv"},
		Fields:     []int{1, 3, 4},
		Delimeter:  "\t",
		Output:     nopWriteCloser{&out},
	}
	if err := Cut(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "a\tc\td\n1\t3\t4\nx\tz\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestCut_OutOfRange(t *testing.T) {
	var out bytes.Buffer
	opts := CutOpts{
		InputFiles: []string{"test_data/data.csv"},
		Fields:     []int{1, 10},
		Delimeter:  "\t",
		Output:     nopWriteCloser{&out},
	}
	if err := Cut(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "a\n1\nx\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestCut_OnlySeparated(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "mixed.txt")
	content := "abc\nx:y:z\n"
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatalf("cannot create file: %v", err)
	}

	var out bytes.Buffer
	opts := CutOpts{
		InputFiles:    []string{file},
		Fields:        []int{1},
		Delimeter:     ":",
		OnlySeparated: true,
		Output:        nopWriteCloser{&out},
	}
	if err := Cut(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "x\n"
	if out.String() != want {
		t.Errorf("got %q, want %q", out.String(), want)
	}
}

func TestCut_NoSuchFile(t *testing.T) {
	var out bytes.Buffer
	opts := CutOpts{
		InputFiles: []string{"not_exist.txt"},
		Fields:     []int{1},
		Delimeter:  "\t",
		Output:     nopWriteCloser{&out},
	}
	err := Cut(opts)
	if err == nil {
		t.Fatal("expected error for non-existent file, got nil")
	}
	if !os.IsNotExist(err) {
		t.Fatalf("expected file not exist error, got %v", err)
	}
}
