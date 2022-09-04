package lcs

import (
	"fmt"

	"github.com/ranpariyachetan/goalgos/utils"
)

// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

func TestMinInsertionForPalindrome() {
	// testMinInsertionForPalindrome("zzazz")
	// testMinInsertionForPalindrome("mbadm")
	// testMinInsertionForPalindrome("leetcode")
	testMinInsertionForPalindrome("aabb")
	testMinInsertionForPalindrome("letelt")

}
func testMinInsertionForPalindrome(s string) {

	rs := utils.ReverseString(s)
	l := len(s)

	memo := make([][]int, l+1)

	for i := range memo {
		memo[i] = make([]int, l+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	lcsLen := findMaxLengthOfLcsWithMemo(s, rs, l, l, memo)

	fmt.Println(l - lcsLen)
}
