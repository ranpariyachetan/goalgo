package knapsack

import (
	"fmt"
	"math"

	"github.com/ranpariyachetan/goalgos/utils"
)

func TestCoinChangeProblem2() {

	//inputs
	// coins := []int{1, 2, 3}
	// sum := 10

	// Following inputs may exceed the time limit. This better be solved by BottomUp approach or with memoization.
	// Answer : 35502874
	// sum := 500
	// coins := []int{3, 5, 7, 8, 9, 10, 11}

	coins := []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}
	sum := 9864

	r := solveCoinChange2TopDown(coins, sum, 0)

	fmt.Println(r)
}

func solveCoinChange2TopDown(coins []int, sum int, n int) int {
	if sum == 0 {
		return 0
	}

	if n >= len(coins) {
		return math.MaxInt - 1
	}

	if coins[n] > sum {
		return solveCoinChange2TopDown(coins, sum, n+1)
	}

	return utils.MinInt(1+solveCoinChange2TopDown(coins, sum-coins[n], n), solveCoinChange2TopDown(coins, sum, n+1))
}

func TestCoinChangeProblem2WithMemo() {

	//inputs
	// coins := []int{1, 2, 3}
	// sum := 10

	// coins := []int{1, 2, 5}
	// sum := 42

	// Following inputs may exceed the time limit. This better be solved by BottomUp approach or with memoization.
	// Answer : 35502874
	// sum := 500
	// coins := []int{3, 5, 7, 8, 9, 10, 11}

	coins := []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}
	sum := 9864

	memo := make([][]int, len(coins)+1)

	for m := range memo {
		memo[m] = make([]int, sum+1)
		for n := range memo[m] {
			memo[m][n] = math.MaxInt - 1
		}
	}

	r := solveCoinChange2TopDownWithMemo(coins, sum, 0, memo)

	fmt.Println(r)
}

func solveCoinChange2TopDownWithMemo(coins []int, sum int, n int, memo [][]int) int {

	if sum == 0 {
		return 0
	}

	if sum < 0 {
		return -1
	}

	if n >= len(coins) || sum < 0 {
		//memo[n][sum] = math.MaxInt - 1
		return math.MaxInt - 1
	}

	// if memo[n][sum] != math.MaxInt-1 {
	// 	return memo[n][sum]
	// }

	// if coins[n] > sum {
	// 	memo[n][sum] = solveCoinChange2TopDownWithMemo(coins, sum, n+1, memo)
	// 	return memo[n][sum]
	// }

	a := solveCoinChange2TopDownWithMemo(coins, sum, n+1, memo)

	b := math.MaxInt

	if coins[n] <= sum {
		b = 1 + solveCoinChange2TopDownWithMemo(coins, sum-coins[n], n, memo)
	}

	//memo[n][sum] = utils.MinInt(1+solveCoinChange2TopDownWithMemo(coins, sum-coins[n], n, memo), solveCoinChange2TopDownWithMemo(coins, sum, n+1, memo))
	//return memo[n][sum]

	return utils.MinInt(a, b)
}
