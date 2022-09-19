package main

import (
	"fmt"
)

func findMedianFromTwoArrays(x, y []int) int {
	i, j := 0, 0
	count := 0
	totalLen := len(x) + len(y)
	s := make([]int, totalLen)
	for {
		if x[i] < y[j] {
			s[count] = x[i]
			i++
		} else {
			s[count] = y[j]
			j++
		}

		count++

		if i >= len(x) {
			for j < len(x) {
				s[count] = y[j]
				count++
				j++
			}
			break
		}

		if j >= len(y) {
			for i < len(x) {
				s[count] = x[i]
				count++
				i++
			}
			break
		}

		// if count >= totalLen {
		// 	break
		// }
	}
	var median int

	if totalLen%2 == 0 {
		median = (s[totalLen/2] + s[(totalLen/2)-1]) / 2

	} else {
		median = s[totalLen/2]
	}

	fmt.Printf("%v\n", s)
	return median
}

func findMedian(x []int) int {
	n := len(x)

	if n == 0 {
		return -1
	}

	if n == 1 {
		return x[0]
	}

	if n%2 == 1 {
		return x[n/2]
	} else {
		return (x[n/2] + x[(n/2)-1]) / 2
	}
}

func testFindMedian() {
	y := []int{3}
	x := []int{4, 7, 9, 11, 13, 15, 18, 20}

	median := findMedianFromTwoArrays(x, y)

	fmt.Println(median)
}

func testMedianFromTwoArrays() {
	x := []int{4, 7, 9, 11, 13, 15, 18}
	y := []int{1, 2, 5, 12, 14}

	median := findMedianFromTwoArrays(x, y)

	fmt.Printf("%d", median)
}
