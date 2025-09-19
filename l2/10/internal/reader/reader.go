package reader

import "unix-sort/internal/args"

// считывание текста из файлов и stdin
func Read(opts args.SortOptions, ch chan []string) {
	if len(opts.Files) == 0 {
		// чтение из STDIN

		// подменяем какой-то ридер?
	}

}
