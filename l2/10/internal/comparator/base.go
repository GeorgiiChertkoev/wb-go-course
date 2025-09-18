package comparator

import (
	"strconv"
	"strings"
)

// базовые компараторы которые не похожи друг на друга
// т.е. -n -M -h и для лексикографический(по умолчанию)

// все функции это Less

func lexicographic(a, b string) bool {
	return a < b
}

func numeric(a, b string) bool {
	// An empty number is treated as ‘0’
	n1, err1 := strconv.ParseFloat(a, 64)
	if err1 != nil {
		n1 = 0
	}
	n2, err2 := strconv.ParseFloat(b, 64)
	if err2 != nil {
		n2 = 0
	}
	if err1 != nil && err2 != nil {
		return a < b // если оба не числа, сортируются лексикографически
	}

	return n1 < n2
}

func month(a, b string) bool {
	months := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

	a = strings.TrimLeft(a, " \t")
	b = strings.TrimLeft(b, " \t")

	var v1, v2 int
	v1, v2 = -1, -1

	for i, m := range months {
		if strings.HasPrefix(a, m) {
			v1 = i
		} else if strings.HasPrefix(b, m) {
			v2 = i
		}
	}
	return v1 < v2
}

func humanNumeric(a, b string) bool {
	a = strings.Trim(a, " \t")
	b = strings.Trim(b, " \t")
	aIsPositive, bIsPositive := true, true
	if len(a) != 0 {
		aIsPositive = a[0] != '-'
	}
	if len(b) != 0 {
		bIsPositive = b[0] != '-'
	}

	if aIsPositive != bIsPositive {
		// - < + -> true
		// false true -> true
		// true false -> false
		return bIsPositive
	}

	getLastChar := func(s string) string {
		// для обработки случая с пустой строкой
		if len(s) != 0 {
			return s[len(s)-1:]
		}
		return "asd" // index по пустой строке возвращает 0, а мы хотим -1
	}

	aScale := strings.Index("KMGTPEZYRQ", getLastChar(a))
	bScale := strings.Index("KMGTPEZYRQ", getLastChar(b))
	if aScale < bScale {
		return true
	} else if aScale > bScale {
		return false
	}

	removeScale := func(a string, scale int) string {
		if scale == -1 {
			return a
		}
		return a[:len(a)-1]
	}
	return numeric(removeScale(a, aScale), removeScale(b, bScale))

}
