package knapsack

import (
	"fmt"
)

func TestKnapSack6BottomUp() {
	arr := []int{1, 1, 2, 3}
	diff := 1

	result := solveKnapsack6BottomUp(arr, diff)

	fmt.Println(result)
}

func solveKnapsack6BottomUp(arr []int, diff int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	subsetSum := (diff + sum) / 2

	subsetSumCount := countSubsetSum(arr, subsetSum)

	return subsetSumCount
}

func countSubsetSum(arr []int, sum int) int {
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
