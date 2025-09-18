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
	pflag.StringVarP(&opts.Separators, "separators", "t", "\t", "разделители для -k")

	// возможные фичи
	// ‘-S size’ ‘--buffer-size=size’
	// ‘-T tempdir’ ‘--temporary-directory=tempdir’
	// ‘--parallel=n’

	pflag.Parse()

	return opts
}
