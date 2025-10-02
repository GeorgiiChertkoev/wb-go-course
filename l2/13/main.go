package main

import (
	"fmt"
	"os"
)

const pipelineChanSize int = 10

func main() {
	// filter (if have onlySeparated flag)
	// split into fields
	// print only chosen fields
	opts, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing args: %v", err)
		return
	}
	err = Cut(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while Cut: %v", err)
	}
}
