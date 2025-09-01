package main

import "fmt"

func DeleteElem[T any](slice []T, i int) []T {
	// 3 разных способа

	// slice = slices.Delete(slice, i, i + 1)

	// slice = append(slice[:i], slice[i+1:]...)

	copy(slice[i:], slice[i+1:]) // перекомпируем элементы со сдвигом на один
	slice = slice[:len(slice)-1] // сокращаем длину

	return slice
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v len:%v cap:%v\n", nums, len(nums), cap(nums))
	nums = DeleteElem(nums, 2)
	fmt.Printf("%v len:%v cap:%v\n", nums, len(nums), cap(nums))
}
