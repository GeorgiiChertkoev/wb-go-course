package main

import (
	"fmt"
	"slices"
)

// сортируем буквы слова
// с помощью map объединяем анаграммы под один ключ
// идем по ключам и переделываем в выходной формат

func GetAnagrams(words []string) map[string][]string {
	tempMap := make(map[string][]string)

	for i := range words {
		w := []rune(words[i])
		slices.Sort(w)
		tempMap[string(w)] = append(tempMap[string(w)], words[i])
	}

	result := make(map[string][]string, len(words))
	for _, v := range tempMap {
		if len(v) <= 1 {
			continue
		}
		result[v[0]] = v
	}
	return result
}

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	for k, v := range GetAnagrams(input) {
		fmt.Printf("%s : %v \n", k, v)
	}

}
