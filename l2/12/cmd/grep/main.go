package main

import (
	"fmt"
	"go-grep/internal/grep"
	"go-grep/internal/options"
	"os"
)

func main() {
	opts, err := options.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse command line arguments: %v\n", err)
		fmt.Fprintf(os.Stderr, "See --help for help\n")
		os.Exit(1)
		return
	}
	results, _ := grep.Grep(*opts)
	fmt.Printf("files: %d\n", len(results))
	for _, greppedFile := range results {
		greppedFile.Print(os.Stdout, "--")
	}
}
