package knapsack

import (
	"fmt"
)

func TestCoinChangeProblem1() {

	//inputs
	coins := []int{1, 2, 3}
	sum := 5

	// Following inputs may exceed the time limit. This better be solved by BottomUp approach or with memoization.
	// Answer : 35502874
	// sum := 500
	// coins := []int{3, 5, 7, 8, 9, 10, 11}

	r := solveCoinChange1TopDown(coins, sum, 0)

	fmt.Println(r)
}

func solveCoinChange1TopDown(coins []int, sum int, n int) int {
	if sum == 0 {
		return 1
	}

	if n >= len(coins) {
		return 0
	}

	x := solveCoinChange1TopDown(coins, sum, n+1)

	if coins[n] <= sum {
		x += solveCoinChange1TopDown(coins, sum-coins[n], n)
	}

	return x
}

func TestCoinChangeProblem1WithMemo() {

	//inputs
	// coins := []int{1, 2, 3}
	// sum := 5

	// Following inputs may exceed the time limit. This better be solved by BottomUp approach or with memoization.
	// Answer : 35502874
	sum := 500
	coins := []int{3, 5, 7, 8, 9, 10, 11}

	memo := make([][]int, len(coins)+1)

	for m := range memo {
		memo[m] = make([]int, sum+1)
		for n := range memo[m] {
			memo[m][n] = -1
		}
	}
	r := solveCoinChange1TopDownWithMemo(coins, sum, 0, memo)

	fmt.Println(r)
}

func solveCoinChange1TopDownWithMemo(coins []int, sum int, n int, memo [][]int) int {
	if sum == 0 {
		return 1
	}

	if n >= len(coins) {
		return 0
	}

	if memo[n][sum] != -1 {
		return memo[n][sum]
	}

	x := solveCoinChange1TopDownWithMemo(coins, sum, n+1, memo)

	if coins[n] <= sum {
		x += solveCoinChange1TopDownWithMemo(coins, sum-coins[n], n, memo)
	}

	memo[n][sum] = x
	return memo[n][sum]
}
