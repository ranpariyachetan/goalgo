package mcm

import (
	"fmt"
	"math"

	"github.com/ranpariyachetan/goalgos/utils"
)

// https://leetcode.com/problems/super-egg-drop/
func TestEggDrop() {
	e := 2
	f := 5

	m := make([][]int, e+1)

	for i := range m {
		m[i] = make([]int, f+1)
		for j := range m[i] {
			m[i][j] = -1
		}
	}
	r := eggDrop(e, f, m)

	fmt.Println(r)
}

func eggDrop(e int, f int, m [][]int) int {
	if f <= 1 {
		return f
	}

	if e == 1 {
		return f
	}

	if m[e][f] != -1 {
		return m[e][f]
	}

	min := math.MaxInt
	for k := 1; k <= f; k++ {
		tmp := 1 + utils.MaxInt(eggDrop(e-1, k-1, m), eggDrop(e, f-k, m))

		if tmp < min {
			min = tmp
		}
	}

	m[e][f] = min
	return m[e][f]
}

func TestEggDropBinary() {
	e := 6
	f := 10000

	m := make([][]int, e+1)

	for i := range m {
		m[i] = make([]int, f+1)
		for j := range m[i] {
			m[i][j] = -1
		}
	}

	r := eggDropBinary(e, f, m)

	fmt.Println(r)
}

func eggDropBinary(k int, n int, m [][]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if k == 1 {
		return n
	}

	if m[k][n] != -1 {
		return m[k][n]
	}

	min := math.MaxInt
	l := 1
	h := n
	tmp := 0

	for l <= h {
		mid := (l + h) / 2

		a := eggDropBinary(k-1, mid-1, m)
		b := eggDropBinary(k, n-mid, m)

		tmp = 1 + utils.MaxInt(a, b)

		if a < b {
			l = mid + 1
		} else {
			h = mid - 1
		}

		min = utils.MinInt(min, tmp)

	}

	m[k][n] = min
	return m[k][n]
}
