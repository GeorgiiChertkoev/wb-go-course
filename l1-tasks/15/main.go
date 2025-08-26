package main

var justString string

//	func someFunc() {
//		v := createHugeString(1 << 10)
//		justString = v[:100]
//	}

// в старом решении мы создали слайс на 1024 байта
// но потом используем только первые 100
// как результат - утечка памяти т.к. сборщик мусора
// не может очистить всю строку так как на часть ее
// ссылается слайс несмотря на то что слайс использует только 100 байт из 1024

func someFunc() {
	// для решения такой проблемы я создаю новый слайс на 100 байт
	// и в него перекопирую только первые 100 символов строки
	// после этого на длинную строку не кто не ссылается и сборщик мусора ее очистит
	v := createHugeString(1 << 10)
	temp := make([]byte, 100)
	copy(temp, v[:100])
	justString = string(temp)
}

func main() {
	someFunc()
}

func createHugeString(length int) string {
	s := make([]byte, length)
	for i := 0; i < length; i++ {
		s[i] = 'a'
	}
	return string(s)
}
