package knapsack

import "fmt"

//https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.
// A subarray is a contiguous part of an array.

func TestFindSubArraysDivByK() {
	arr := []int{4, 5, 0, -2, -3, 1}
	k := 5

	r := findSubArraysDivByK(arr, k)

	fmt.Println(r)
}

func findSubArraysDivByK(nums []int, k int) int {
	n := len(nums) - 1
	return getCountOfSubArrays(nums, nums[n], k, n)
}

func getCountOfSubArrays(nums []int, sum int, k int, n int) int {
	if sum%k == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if nums[n]%k == 0 {
		return 1 + getCountOfSubArrays(nums, nums[n]+nums[n-1], k, n-1)
	} else {
		return getCountOfSubArrays(nums, nums[n]+nums[n-1], k, n-1)
	}
}
