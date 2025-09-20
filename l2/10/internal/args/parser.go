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

func ParseArgs() SortOptions {
	var opts SortOptions

	pflag.IntVarP(&opts.KeyColumn, "key", "k", 0, "номер колонки для сортировки (по умолчанию 0 = вся строка)")
	pflag.BoolVarP(&opts.Numeric, "numeric-sort", "n", false, "сортировка по числовому значению")
	pflag.BoolVarP(&opts.Reverse, "reverse", "r", false, "обратный порядок")
	pflag.BoolVarP(&opts.Unique, "unique", "u", false, "выводить только уникальные строки")
	pflag.BoolVarP(&opts.Month, "month-sort", "M", false, "сортировка по названию месяца")
	pflag.BoolVarP(&opts.IgnoreBlanks, "ignore-leading-blanks", "b", false, "игнорировать ведущие пробелы") // соответствует документации gnu
	pflag.BoolVarP(&opts.CheckSorted, "check", "c", false, "проверить, отсортированы ли данные")
	pflag.BoolVarP(&opts.HumanNumeric, "human-numeric-sort", "h", false, "сортировка с учётом суффиксов (K, M, G)")

	// доп фичи
	pflag.StringVarP(&opts.Separators, "field-separator", "t", "\t", "разделители для -k")
	pflag.StringVarP(&opts.TempDir, "temporary-directory", "T", "tmp", "директория для хранения временных файлов")
	pflag.StringVarP(&opts.OutputFile, "output", "o", "", "файл в который пишутся отсортированные строки")
	pflag.Int64VarP(&opts.BufSize, "buffer-size", "S", 256*1024*1024, "размер буффера при считывании данных") // 256 мб по умолчанию

	// ‘--parallel=n’

	pflag.Parse()

	opts.Files = pflag.Args()

	return opts
}
