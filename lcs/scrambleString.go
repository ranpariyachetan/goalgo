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

func TestScrambleStringWithMemo() {
	s1 := "ccabcbabcbabbbbcbb"
	s2 := "bbbbabccccbbbabcba"
	m := make(map[string]bool)
	r := checkIfScrambleWithMemo(s1, s2, m)
	fmt.Println(r)
}

func checkIfScrambleWithMemo(s1 string, s2 string, m map[string]bool) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) <= 1 || len(s2) <= 1 {
		return false
	}

	if _, ok := m[s1+":"+s2]; ok {
		return m[s1+":"+s2]
	}

	for i := 0; i < len(s1)-1; i++ {
		b1 := checkIfScrambleWithMemo(s1[0:i+1], s2[len(s2)-(i+1):], m) && checkIfScrambleWithMemo(s1[i+1:], s2[0:len(s2)-(i+1)], m)
		b2 := checkIfScrambleWithMemo(s1[0:i+1], s2[0:i+1], m) && checkIfScrambleWithMemo(s1[i+1:], s2[i+1:], m)

		if b1 || b2 {
			m[s1+":"+s2] = true
			return true
		}
	}

	m[s1+":"+s2] = false
	return false
}
