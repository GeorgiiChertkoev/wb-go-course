package comparator

import (
	"strings"
)

// флаги которые меняют входную строку
// т.е. -k, -b

func getCollumn(s string, collumnId int, separators string) string {
	columns := strings.FieldsFunc(s, func(letter rune) bool {
		return strings.ContainsRune(separators, letter)
	})
	// счет с одного
	if len(columns) >= collumnId {
		return columns[collumnId-1]
	}
	return ""
}

func ignoreBlanks(s string) string {
	return strings.Trim(s, " \t")
}
