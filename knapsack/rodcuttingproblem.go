package knapsack

import (
	"fmt"

	"github.com/ranpariyachetan/goalgos/utils"
)

// Given a rod of length ‘N’ units. The rod can be cut into different sizes and each size has a cost associated with it.
// Determine the maximum cost obtained by cutting the rod and selling its pieces.

func TestRodCuttingProblem() {
	maxLen := 8

	costs := []int{1, 5, 8, 9, 10, 17, 17, 20}

	lens := make([]int, 8)

	for i := 0; i < maxLen; i++ {
		lens[i] = i + 1
	}

	answer := solveRoadCuttingTopDown(costs, lens, maxLen)

	fmt.Println(answer)
}

func solveRoadCuttingTopDown(costs []int, lens []int, maxLen int) int {
	n := len(costs)

	return rodCuttingTopDown(costs, lens, maxLen, n)
}

func rodCuttingTopDown(costs []int, lens []int, maxLen int, n int) int {
	if n == 0 || maxLen == 0 {
		return 0
	}

	if lens[n-1] > maxLen {
		return knapsack1TopDown(costs, lens, maxLen, n-1)
	}

	a := costs[n-1] + knapsack1TopDown(costs, lens, maxLen-lens[n-1], n)
	b := knapsack1TopDown(costs, lens, maxLen, n-1)

	return utils.MaxInt(a, b)
}
