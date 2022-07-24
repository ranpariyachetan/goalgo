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

	if vals[n] == sum {
		return true
	}

	if vals[n] < sum {
		return knapsack2(vals, sum-vals[n], n+1) || knapsack2(vals, sum, n+1)
	} else {
		return knapsack2(vals, sum, n+1)
	}
}
