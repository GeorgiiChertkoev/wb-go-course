package main

import (
	"fmt"
	"unicode"
)

func futureStringLen(s string) int {
	// считает длину распакованной строки учитывая escape sequence
	// и распоковку нулевой длины
	curNum := 0
	resultLen := 0
	numsInRow := 0
	skipNext := false
	for _, v := range s {
		resultLen++
		if skipNext {
			skipNext = false
			continue
		}
		if v == rune('\\') {
			resultLen--
			skipNext = true
		}
		if unicode.IsDigit(v) {
			curNum = curNum*10 + int(v-'0')
			numsInRow++
		} else if curNum != 0 {
			resultLen += curNum - numsInRow - 1
			numsInRow = 0
			curNum = 0
		}
	}
	if curNum != 0 {
		resultLen += curNum - numsInRow - 1
	}
	return resultLen
}

func unpackIntoSlice(dest []rune, letter rune, n int) []rune {
	for _ = range n {
		dest = append(dest, letter)
	}
	return dest
}

func Unpack(packed string) (string, error) {
	if packed == "" {
		return "", nil
	}
	if unicode.IsDigit(rune(packed[0])) {
		return "", fmt.Errorf("first cymbol is number, nothing to unpack")
	}
	curNum := 0
	hasNums := false // нужно чтобы различать случаи когда не было цифр и когда был 0
	escapeMode := false
	var prevLetter rune

	unpacked := make([]rune, 0, futureStringLen(packed))
	for _, v := range packed {
		if v == rune('\\') {
			escapeMode = true
		} else if unicode.IsDigit(v) && !escapeMode {
			curNum = curNum*10 + int(v-'0')
			hasNums = true
		} else {
			if !hasNums {
				curNum = 1
			}
			if prevLetter != 0 {
				unpacked = unpackIntoSlice(unpacked, prevLetter, curNum)
			}
			curNum = 0
			hasNums = false
			escapeMode = false

			prevLetter = v
		}
	}
	if !hasNums {
		curNum = 1
	}
	unpacked = unpackIntoSlice(unpacked, prevLetter, curNum)

	return string(unpacked), nil
}

func main() {
	fmt.Println(futureStringLen(`\03\a2b2\`))
	fmt.Println(Unpack(`\03\a\2b2\`))
}
