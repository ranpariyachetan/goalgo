package knapsack

import "fmt"

/*Given weights and values of n items, put these items in a knapsack of capacity W to get the maximum total value in the knapsack.
In other words, given two integer arrays val[0..n-1] and wt[0..n-1] which represent values and weights associated with n items respectively.
Also given an integer W which represents knapsack capacity, find out the maximum value subset of val[] such that
sum of the weights of this subset is smaller than or equal to W. You cannot break an item, either pick the complete item or don’t pick it (0-1 property). */

func TestKnapSack1() {
	vals := []int{60, 100, 120}
	weights := []int{10, 20, 30}

	w := 50

	answer := solveKnapsack1(vals, weights, w)

	fmt.Println(answer)
}
func solveKnapsack1(vals []int, weights []int, capacity int) int {
	n := len(vals)

	return knapsack1(vals, weights, capacity, n)
}

func knapsack1(vals []int, weights []int, capacity int, n int) int {
	if n == 0 || capacity == 0 {
		return 0
	}

	if weights[n-1] > capacity {
		knapsack1(vals, weights, capacity, n-1)
	}
	//if weights[n-1] <= capacity {
	a := vals[n-1] + knapsack1(vals, weights, capacity-weights[n-1], n-1)
	b := knapsack1(vals, weights, capacity, n-1)

	if a > b {
		return a
	} else {
		return b
	}
}
