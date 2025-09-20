package args

import (
	"fmt"

	"github.com/spf13/pflag"
)

type SortOptions struct {
	KeyColumn    int
	Numeric      bool
	Reverse      bool
	Unique       bool
	Month        bool
	IgnoreBlanks bool
	CheckSorted  bool
	HumanNumeric bool
	Separators   string
	BufSize      int64  // size of buffer for reading in bytes
	TempDir      string // dir for temp files
	KeepTemps    bool
	OutputFile   string
	Files        []string
}

func (opts SortOptions) Print() {
	fmt.Printf("KeyColumn: %d\n", opts.KeyColumn)
	fmt.Printf("Numeric: %v\n", opts.Numeric)
	fmt.Printf("Reverse: %v\n", opts.Reverse)
	fmt.Printf("Unique: %v\n", opts.Unique)
	fmt.Printf("Month: %v\n", opts.Month)
	fmt.Printf("IgnoreBlanks: %v\n", opts.IgnoreBlanks)
	fmt.Printf("CheckSorted: %v\n", opts.CheckSorted)
	fmt.Printf("HumanNumeric: %v\n", opts.HumanNumeric)
	fmt.Printf("Separetor: %q\n", opts.Separators)
}

func ParseArgs(args []string) SortOptions {
	var opts SortOptions

	fs := pflag.NewFlagSet("sort", pflag.ContinueOnError)

	fs.IntVarP(&opts.KeyColumn, "key", "k", 0, "номер колонки для сортировки (по умолчанию 0 = вся строка)")
	fs.BoolVarP(&opts.Numeric, "numeric-sort", "n", false, "сортировка по числовому значению")
	fs.BoolVarP(&opts.Reverse, "reverse", "r", false, "обратный порядок")
	fs.BoolVarP(&opts.Unique, "unique", "u", false, "выводить только уникальные строки")
	fs.BoolVarP(&opts.Month, "month-sort", "M", false, "сортировка по названию месяца")
	fs.BoolVarP(&opts.IgnoreBlanks, "ignore-leading-blanks", "b", false, "игнорировать ведущие пробелы") // соответствует документации gnu
	fs.BoolVarP(&opts.CheckSorted, "check", "c", false, "проверить, отсортированы ли данные")
	fs.BoolVarP(&opts.HumanNumeric, "human-numeric-sort", "h", false, "сортировка с учётом суффиксов (K, M, G)")

	// доп фичи
	fs.StringVarP(&opts.Separators, "field-separator", "t", "\t", "разделители для -k")
	fs.StringVarP(&opts.TempDir, "temporary-directory", "T", "", "директория для хранения временных файлов")
	fs.StringVarP(&opts.OutputFile, "output", "o", "", "файл в который пишутся отсортированные строки")
	fs.Int64VarP(&opts.BufSize, "buffer-size", "S", 256*1024*1024, "размер буффера при считывании данных") // 256 мб по умолчанию
	fs.BoolVar(&opts.KeepTemps, "keep-temps", false, "не удалять временные файлы после сортировки")

	fs.Parse(args)

	opts.Files = fs.Args()

	return opts
}
