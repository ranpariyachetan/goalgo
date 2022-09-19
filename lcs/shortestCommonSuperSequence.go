package lcs

import (
	"fmt"

	"github.com/ranpariyachetan/goalgos/utils"
)

func TestFindShortestCommonSuperSequence() {
	f := findShortestCommonSuperSequenceBottomUp("abcdaf", "acbcf")

	fmt.Println(f)
}
func findShortestCommonSuperSequence(x string, y string, m int, n int, o *string) string {
	if m == 0 || n == 0 {
		return ""
	}

	// if m == 0 {
	// 	*o += string(y[n-1])
	// 	return string(y[n-1]) + findShortestCommonSuperSequence(x, y, m, n-1, o)
	// }

	// if n == 0 {
	// 	*o += string(x[m-1])
	// 	return string(x[m-1]) + findShortestCommonSuperSequence(x, y, m-1, n, o)
	// }

	if x[m-1] == y[n-1] {
		*o += string(x[m-1])
		return string(x[m-1]) + findShortestCommonSuperSequence(x, y, m-1, n-1, o)
	} else {
		// *o += string(x[m-1]) + string(y[n-1])
		return findShortestCommonSuperSequence(x, y, m, n-1, o) + findShortestCommonSuperSequence(x, y, m-1, n, o)
	}
}

func findShortestCommonSuperSequenceBottomUp(x string, y string) string {

	m := len(x)
	n := len(y)

	t := make([][]int, m+1)

	for i := range t {
		t[i] = make([]int, n+1)
		for j := range t[i] {
			if i == 0 || j == 0 {
				t[i][j] = 0
			}
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = utils.MaxInt(t[i-1][j], t[i][j-1])
			}
		}
	}

	i := m
	j := n

	s := ""
	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			s = string(x[i-1]) + s
			i--
			j--
		} else {
			if t[i][j-1] > t[i-1][j] {
				s = string(y[j-1]) + s
				j--
			} else {
				s = string(x[i-1]) + s
				i--
			}
		}
	}

	for i > 0 {
		s = string(x[i-1]) + s
		i--
	}

	for j > 0 {
		s = string(y[j-1]) + s
		j--
	}
	return s
}
