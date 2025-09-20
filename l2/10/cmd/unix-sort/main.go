package main

import (
	"fmt"
	"os"
	"unix-sort/internal/args"
	"unix-sort/internal/sorter"
)

func main() {
	err := sorter.Sort(args.ParseArgs(os.Args))
	if err != nil {
		fmt.Println(err)
	}
}
