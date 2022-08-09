package lcs

import (
	"fmt"

	"github.com/ranpariyachetan/goalgos/utils"
)

func TestFindMaxLengthOfLcs() {
	t1 := "abcde"
	t2 := "ace"

	r := findMaxLengthOfLsc(t1, t2, len(t1), len(t2))

	fmt.Println(r)
}
func findMaxLengthOfLsc(x string, y string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if x[m] == y[m] {
		return 1 + findMaxLengthOfLsc(x, y, m-1, n-1)
	} else {
		return utils.MaxInt(findMaxLengthOfLsc(x, y, m, n-1), findMaxLengthOfLsc(x, y, m-1, n))
	}
}

func TestFindMaxLengthOfLcsWithMemo() {
	t1 := "mhunuzqrkzsnidwbun"
	t2 := "szulspmhwpazoxijwbq"

	m := len(t1)
	n := len(t2)

	memo := make([][]int, m+1)

	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	r := findMaxLengthOfLcsWithMemo(t1, t2, m, n, memo)

	fmt.Println(r)
}

func findMaxLengthOfLcsWithMemo(x string, y string, m int, n int, memo [][]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if memo[m][n] != -1 {
		return memo[m][n]
	}

	if x[m] == y[m] {
		memo[m][n] = 1 + findMaxLengthOfLsc(x, y, m-1, n-1)
		return memo[m][n]
	} else {
		memo[m][n] = utils.MaxInt(findMaxLengthOfLsc(x, y, m, n-1), findMaxLengthOfLsc(x, y, m-1, n))
		return memo[m][n]
	}
}
