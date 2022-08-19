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

	r := eggDrop(e, f)

	fmt.Println(r)
}

func eggDrop(e int, f int) int {
	if f <= 1 {
		return f
	}

	if e == 1 {
		return f
	}

	min := math.MaxInt
	for k := 1; k <= f; k++ {
		tmp := 1 + utils.MaxInt(eggDrop(e-1, k-1), eggDrop(e, f-k))

		if tmp < min {
			min = tmp
		}
	}

	return min
}
