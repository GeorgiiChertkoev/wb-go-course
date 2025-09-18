package comparator

import (
	"fmt"
	"unix-sort/internal/args"
)

// объединяет компараторы и декораторы и возвращает функцию сравнения

func BuildComparator(options args.SortOptions) (func(a, b string) bool, error) {
	sufficientBases := 0 // numeric month humanNumeric  если больше одного, то ошибка из-за неоднозначности
	base := lexicographic
	if options.Numeric {
		base = numeric
		sufficientBases++
	}
	if options.HumanNumeric {
		base = humanNumeric
		sufficientBases++
	}
	if options.Month {
		base = month
		sufficientBases++
	}
	if sufficientBases > 1 {
		return nil, fmt.Errorf("incompatible sorting parameters")
	}
	currFunc := base
	if options.KeyColumn != 0 {
		currFunc = composeFunctions(func(s string) string {
			return getCollumn(s, options.KeyColumn, options.Separators)
		}, base)
	}
	if options.IgnoreBlanks {
		currFunc = composeFunctions(ignoreBlanks, currFunc)
	}
	if options.Reverse {
		return func(a, b string) bool {
			return !currFunc(a, b)
		}, nil
	}
	return currFunc, nil
}

func composeFunctions(decorator func(s string) string, base func(a, b string) bool) func(a, b string) bool {
	return func(a, b string) bool {
		return base(decorator(a), decorator(b))
	}
}
