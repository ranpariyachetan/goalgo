package knapsack

import (
	"fmt"

	"github.com/ranpariyachetan/goalgos/utils"
)

/*Given weights and values of n items, put these items in a knapsack of capacity W to get the maximum total value in the knapsack.
In other words, given two integer arrays val[0..n-1] and wt[0..n-1] which represent values and weights associated with n items respectively.
Also given an integer W which represents knapsack capacity, find out the maximum value subset of val[] such that
sum of the weights of this subset is smaller than or equal to W. You cannot break an item, either pick the complete item or don’t pick it (0-1 property). */

func TestKnapSack1TopDown() {
	vals := []int{60, 100, 120}
	weights := []int{10, 20, 30}

	w := 50

	answer := solveKnapsack1TopDown(vals, weights, w)

	fmt.Println(answer)
}

func solveKnapsack1TopDown(vals []int, weights []int, capacity int) int {
	n := len(vals)

	return knapsack1TopDown(vals, weights, capacity, n)
}

func knapsack1TopDown(vals []int, weights []int, capacity int, n int) int {
	if n == 0 || capacity == 0 {
		return 0
	}

	if weights[n-1] > capacity {
		return knapsack1TopDown(vals, weights, capacity, n-1)
	}

	a := vals[n-1] + knapsack1TopDown(vals, weights, capacity-weights[n-1], n-1)
	b := knapsack1TopDown(vals, weights, capacity, n-1)

	return utils.MaxInt(a, b)
}

func TestKnapSack1TopDownMemo() {
	vals := []int{60, 100, 120}
	weights := []int{10, 20, 30}

	w := 50

	answer := solveKnapsack1WithTopDownMemo(vals, weights, w)

	fmt.Println(answer)
}

func solveKnapsack1WithTopDownMemo(vals []int, weights []int, w int) int {
	n := len(vals)

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			dp[i][j] = -1
		}
	}

	return knapsack1WithTopDownMemo(vals, weights, w, n, dp)
}

func knapsack1WithTopDownMemo(vals []int, weights []int, w int, n int, dp [][]int) int {

	if w == 0 || n == 0 {
		return 0
	}

	if dp[n][w] != -1 {
		return dp[n][w]
	}

	if weights[n-1] <= w {
		dp[n][w] = utils.MaxInt(vals[n-1]+knapsack1TopDown(vals, weights, w-weights[n-1], n-1), knapsack1TopDown(vals, weights, w, n-1))
		return dp[n][w]
	} else {
		dp[n][w] = knapsack1TopDown(vals, weights, w, n-1)
		return dp[n][w]
	}
}

func TestKnapSack1BottomUp() {
	vals := []int{60, 100, 120}
	weights := []int{10, 20, 30}

	w := 50

	answer := solveKnapsack1BottomUp(vals, weights, w)

	fmt.Println(answer)
}

func solveKnapsack1BottomUp(vals []int, weigths []int, w int) int {
	n := len(vals)
	return knapsack1BottomUp(vals, weigths, w, n)
}

func knapsack1BottomUp(vals []int, weights []int, w int, n int) int {

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < w+1; j++ {
			if weights[i-1] <= j {
				dp[i][j] = utils.MaxInt(vals[i-1]+dp[i-1][j-weights[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][w]
}
