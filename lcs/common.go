package lcs

import "github.com/ranpariyachetan/goalgos/utils"

func findMaxLengthOfLcsWithMemo(x string, y string, m int, n int, memo [][]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if memo[m][n] != -1 {
		return memo[m][n]
	}

	if x[m-1] == y[n-1] {
		memo[m][n] = 1 + findMaxLengthOfLcsWithMemo(x, y, m-1, n-1, memo)
		return memo[m][n]
	} else {
		memo[m][n] = utils.MaxInt(findMaxLengthOfLcsWithMemo(x, y, m, n-1, memo), findMaxLengthOfLcsWithMemo(x, y, m-1, n, memo))
		return memo[m][n]
	}
}
