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
		return
	}
	results, err := grep.Grep(*opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while grepping: %v", err)
		return
	}
	fmt.Printf("files: %d\n", len(results))
	for i, greppedFile := range results {
		if len(results) > 1 {
			fmt.Printf("File: %s\n", opts.Files[i])
		}
		greppedFile.Print(os.Stdout)
	}
}
