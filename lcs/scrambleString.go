package lcs

import "fmt"

// https://leetcode.com/problems/scramble-string/

func TestIsScramble() {

	s1 := "ccabcbabcbabbbbcbb"
	s2 := "bbbbabccccbbbabcba"
	r := checkIfScramble(s1, s2)
	fmt.Println(r)
}

func checkIfScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) <= 1 || len(s2) <= 1 {
		return false
	}

	for i := 0; i < len(s1)-1; i++ {
		b1 := checkIfScramble(s1[0:i+1], s2[len(s2)-(i+1):]) && checkIfScramble(s1[i+1:], s2[0:len(s2)-(i+1)])
		b2 := checkIfScramble(s1[0:i+1], s2[0:i+1]) && checkIfScramble(s1[i+1:], s2[i+1:])

		if b1 || b2 {
			return true
		}
	}
	return false
}
