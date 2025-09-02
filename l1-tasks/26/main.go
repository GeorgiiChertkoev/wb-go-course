package main

import (
	"fmt"
	"unicode"
)

func HasDuplicates(s string) bool {
	// если предположить что строка не содержит рун или у нас совсем небольшой набор возможных символов
	// то можно было бы просто создать массив на кол-во букв и тогда проверка и добавление было бы
	// за настоящую О(1) а не амортизированную
	// но я реализовал чтобы работало для любых рун

	m := make(map[rune]struct{}) // создаем множество
	for _, letter := range s {
		letter = unicode.ToLower(letter) // чтобы не различать буквы в разных регистрах
		if _, ok := m[letter]; ok {
			return true
		}
		m[letter] = struct{}{}
	}
	return false
}

func main() {
	fmt.Printf("HasDuplicates('abcd'):	       %t\n", HasDuplicates("abcd"))
	fmt.Printf("HasDuplicates('abCdefAaf'):    %t\n", HasDuplicates("abCdefAaf"))
	fmt.Printf("HasDuplicates('AbCd Норм'):    %t\n", HasDuplicates("AbCd Норм"))
	fmt.Printf("HasDuplicates('aabcd'):        %t\n", HasDuplicates("aabcd"))
	fmt.Printf("HasDuplicates('Привет Tебе'):  %t\n", HasDuplicates("Привет тебе"))
}
