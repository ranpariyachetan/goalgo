package knapsack

import "fmt"

//Given a set of non-negative integers, and a value sum, determine if there is a subset of the given set with sum equal to given sum.

func TestKnapSack2() {
	vals := []int{5, 10, 15, 20}
	sum := 20

	result := solveKnapsack2(vals, sum)

	fmt.Println(result)
}

func solveKnapsack2(vals []int, sum int) bool {

	return knapsack2(vals, sum, 0)
}

func knapsack2(vals []int, sum int, n int) bool {

	if sum == 0 {
		return true
	}
	if n >= len(vals) {
		return false
	}

	if vals[n] <= sum {
		return knapsack2(vals, sum-vals[n], n+1) || knapsack2(vals, sum, n+1)
	} else {
		return knapsack2(vals, sum, n+1)
	}
}

func TestKnapSack2WithDP() {
	vals := []int{5, 10, 15, 20}
	sum := 25

	result := solveKnapsack2WithDP(vals, sum)

	fmt.Println(result)
}

func solveKnapsack2WithDP(vals []int, sum int) bool {
	n := len(vals)
	dp := make([][]bool, n+1, sum+1)

	for i := range dp {
		dp[i] = make([]bool, sum+1)
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = false
			}

			if j == 0 {
				dp[i][j] = true
			}
		}
	}

	return knapsack2WithDP(vals, sum, dp)
}

func knapsack2WithDP(vals []int, sum int, dp [][]bool) bool {
	n := len(vals)

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if vals[i-1] <= j {
				dp[i][j] = dp[i][j-vals[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][sum]
}
