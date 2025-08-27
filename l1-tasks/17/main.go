package main

import (
	"fmt"
	"slices"
)

func BinarySearch(nums []int, v int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] > v {
			r = m
		}
		if nums[m] < v {
			l = m
		}
		if nums[m] == v {
			return m
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 4, 5, 7, 8, 11, 34, 124}
	slices.Sort(nums)

	fmt.Printf("%d on pos %d in %v\n", 34, BinarySearch(nums, 34), nums)
}
