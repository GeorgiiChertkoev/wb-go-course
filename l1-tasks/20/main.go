package main

import (
	"fmt"
	"slices"
	"strings"
	"unsafe"
)

func SimplestSolution(s string) string {
	words := strings.Split(s, " ")
	slices.Reverse(words)
	return strings.Join(words, " ")
}

func WithOneCopy(s string) string {
	reversed := make([]byte, len(s))

	l := 0 // начало нынешнего слова
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for j := l; j < i; j++ {
				reversed[len(s)-i+j-l] = s[j]
			}
			if l != 0 { // добавляем пробел для не первого слова
				reversed[len(s)-l] = ' '
			}
			l = i + 1
		}
	}
	// добавляем последнее слово в начало
	for j := 0; j < len(s)-l; j++ {
		reversed[j] = s[l+j]
	}
	reversed[len(s)-l] = ' '

	return unsafe.String(&reversed[0], len(reversed)) // избегаем ненужного копирования напирямую отдавая наш массив в строку
}

func main() {
	sentance := string("прив как дела й окак")

	fmt.Printf("was: \"%s\" \n", sentance)
	fmt.Printf("got: \"%s\" \n", SimplestSolution(sentance))
	fmt.Printf("got: \"%s\" \n", WithOneCopy(sentance))
}
