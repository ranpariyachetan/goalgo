package arrays

import "fmt"

func TestBinarySearch() {
	// arr := []int{-1, 0, 3, 5, 9, 12}

	arr := []int{2}

	target := 2

	r := search(arr, target)

	fmt.Println(r)
}
func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l int, h int, target int) int {
	if l >= h {
		return -1
	}

	m := (l + h) / 2

	if target == nums[m] {
		return m
	}

	if target < nums[m] {
		return binarySearch(nums, l, m, target)
	}

	return binarySearch(nums, m+1, h, target)
}
