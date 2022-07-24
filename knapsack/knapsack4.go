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

func TestKnapSack4BottomUp() {
	arr := []int{6, 9, 7, 8, 8}
	sum := 15

	// arr := []int{1, 1, 1, 1}
	// sum := 1

	result := knapsack4BottomUp(arr, sum)

	fmt.Println(result)
}

func knapsack4BottomUp(arr []int, sum int) int {
	n := len(arr)

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = 0
			}
			if j == 0 {
				dp[i][j] = 1

			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][sum]
}
