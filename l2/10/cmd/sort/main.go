package main

import (
	"fmt"
	"strings"
	"unix-sort/internal/args"

	"github.com/spf13/pflag"
)

func main() {
	opts := args.ParseArgs()
	opts.Print()

	files := pflag.Args()
	if len(files) == 0 {
		fmt.Println("Читаем из stdin")
	} else {
		fmt.Printf("Файлы: %s\n", strings.Join(files, ", "))
	}

}
