package knapsack

import "fmt"

// Equal Sum Partition Problem.
// Given an array arr[]. Determine whether it is possible to split the array into two sets such that the sum of elements in both the sets is equal.
// If it is possible, then return true else return false.

func TestKnapSack3TopDown() {
	arr := []int{6, 9, 7, 8}

	result := solveKnapsack3TopDown(arr)

	fmt.Println(result)
}

func solveKnapsack3TopDown(arr []int) bool {

	sum := 0

	for _, num := range arr {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	return knapsack3TopDown(arr, sum, 0)
}

func knapsack3TopDown(arr []int, sum int, n int) bool {
	if sum == 0 {
		return true
	}

	if n >= len(arr) {
		return false
	}

	if arr[n] <= sum {
		return knapsack3TopDown(arr, sum-arr[n], n+1) || knapsack3TopDown(arr, sum, n+1)
	} else {
		return knapsack3TopDown(arr, sum, n+1)
	}
}
