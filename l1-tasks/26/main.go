package main

import (
	"fmt"
	"unicode"
)

func AllUnique(s string) bool {
	// если предположить что строка не содержит рун или у нас совсем небольшой набор возможных символов
	// то можно было бы просто создать массив на кол-во букв и тогда проверка и добавление было бы
	// за настоящую О(1) а не амортизированную
	// но я реализовал чтобы работало для любых рун

	m := make(map[rune]struct{}) // создаем множество
	for _, letter := range s {
		letter = unicode.ToLower(letter) // чтобы не различать буквы в разных регистрах
		if _, ok := m[letter]; ok {
			return false
		}
		m[letter] = struct{}{}
	}
	return true
}

func main() {
	fmt.Printf("AllUnique('abcd'):         %t\n", AllUnique("abcd"))
	fmt.Printf("AllUnique('abCdefAaf'):    %t\n", AllUnique("abCdefAaf"))
	fmt.Printf("AllUnique('AbCd Норм'):    %t\n", AllUnique("AbCd Норм"))
	fmt.Printf("AllUnique('aabcd'):        %t\n", AllUnique("aabcd"))
	fmt.Printf("AllUnique('Привет Tебе'):  %t\n", AllUnique("Привет тебе"))
}
