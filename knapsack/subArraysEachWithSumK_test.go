package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	arr    []int
	target int
	want   int
}

func TestSubArraysWithEachSumK(t *testing.T) {

	cases := []testCase{
		{arr: []int{3, 2, 2, 4, 3}, target: 3, want: 2},
		{arr: []int{7, 3, 4, 7}, target: 7, want: 2},
		{arr: []int{4, 3, 2, 6, 2, 3, 4}, target: 6, want: -1},
		{arr: []int{1, 2, 2, 3, 2, 6, 7, 2, 1, 4, 6, 8}, target: 5, want: 3},
	}

	for _, tc := range cases {

		r := minSumOfLengths(tc.arr, tc.target)

		assert.Equal(t, tc.want, r)
	}
}
