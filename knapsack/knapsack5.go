package knapsack

import (
	"fmt"
	"math"

	"github.com/ranpariyachetan/goalgos/utils"
)

func TestKnapSack5BottomUp() {
	arr := []int{1, 2, 7}

	result := solveKnapsack5BottomUp(arr)

	fmt.Println(result)
}

func solveKnapsack5BottomUp(arr []int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	res := subsetBottomUp(arr, sum)

	minDiff := math.MaxInt

	for i, item := range res[:len(res)/2] {
		if item {
			minDiff = utils.MinInt(minDiff, sum-(2*i))
		}
	}

	return minDiff
}
func subsetBottomUp(arr []int, sum int) []bool {
	n := len(arr)

	dp := make([][]bool, n+1)

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

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n]
}
