package knapsack

import (
	"fmt"
	"sort"
)

func minSumOfLengths(arr []int, target int) int {

	counts := make([]int, 0)

	sum := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			sum = 0
			count = 0
			counts = append(counts, 1)
		} else if arr[i] < target {
			sum += arr[i]
			count++

			if sum == target {
				sum = 0
				counts = append(counts, count)
				count = 0
			} else if sum > target {
				sum = 0
				count = 0
			}
		}
	}

	if len(counts) < 2 {
		return -1
	}

	sort.Ints(counts)

	fmt.Println(counts)

	return counts[0] + counts[1]
}
