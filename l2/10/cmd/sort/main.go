package main

import (
	"fmt"
	"unix-sort/internal/args"
	"unix-sort/internal/sorter"
)

func main() {
	err := sorter.Sort(args.ParseArgs())
	if err != nil {
		fmt.Println(err)
	}
}
