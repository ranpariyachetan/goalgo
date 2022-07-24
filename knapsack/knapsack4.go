package knapsack

import "fmt"

// Given an array arr[] of length N and an integer X, the task is to find the number of subsets with a sum equal to X.

func TestKnapSack4TopDown() {
	// arr := []int{6, 9, 7, 8, 8}
	// sum := 15

	arr := []int{1, 1, 1, 1}
	sum := 1

	result := solveKnapsack4TopDown(arr, sum)

	fmt.Println(result)
}

func solveKnapsack4TopDown(arr []int, sum int) int {
	return knapsack4TopDown(arr, sum, len(arr)-1)
}

func knapsack4TopDown(arr []int, sum int, n int) int {
	if sum == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if arr[n] <= sum {
		return knapsack4TopDown(arr, sum-arr[n], n-1) + knapsack4TopDown(arr, sum, n-1)
	} else {
		return knapsack4TopDown(arr, sum, n-1)
	}
}
