package main

import "fmt"

func partition(nums []int, l int, r int) int {
	pivot := nums[(l+r)/2]
	for l < r {
		for nums[l] < pivot {
			l++
		}
		for nums[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		l++
		r--
		nums[l], nums[r] = nums[r], nums[l]
	}
	return r
}

func quickSortInplaceImpl(nums []int, l int, r int) {
	if l < r {
		pviotPos := partition(nums, l, r)
		quickSortInplaceImpl(nums, l, pviotPos)
		quickSortInplaceImpl(nums, pviotPos+1, r)
	}
}

func QuickSortInplace(nums []int) {
	quickSortInplaceImpl(nums, 0, len(nums)-1)
}

func QuickSort(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	QuickSortInplace(sorted)
	return sorted
}

func main() {
	// nums := []int{14, 5, 5, 5, 5, 1, 45, 7, 2, 4}
	nums := []int{5, 5, 14, 45, 7, 5, 5}
	sorted := QuickSort(nums)
	fmt.Printf("was: %v\n", nums)
	fmt.Printf("got: %v\n", sorted)
}
