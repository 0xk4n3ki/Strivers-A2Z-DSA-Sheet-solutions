package main

import (
	"fmt"
)

// [OLD] https://takeuforward.org/data-structure/median-of-row-wise-sorted-matrix/
// [OLD] https://takeuforward.org/data-structure/peak-element-in-array/
// [OLD] https://takeuforward.org/data-structure/count-reverse-pairs/
// [NEW] https://takeuforward.org/arrays/kth-missing-positive-number/
// [NEW] https://leetcode.com/problems/string-to-integer-atoi/description/
// [NEW] https://takeuforward.org/data-structure/power-set-print-all-the-possible-subsequences-of-the-string/
// [NEW] https://takeuforward.org/data-structure/word-search-leetcode/
// [NEW] https://leetcode.com/problems/subsets/description/
// [NEW] https://takeuforward.org/data-structure/check-for-balanced-parentheses/
// [NEW] https://leetcode.com/problems/remove-k-digits/
// [NEW] https://leetcode.com/problems/longest-repeating-character-replacement/description/

func main() {
	s := "AABABBBCCD"
	k := 2
	ans := longest(s, k)
	fmt.Printf("ans: %d\n", ans)
}

func longest(s string, k int) int {
	maxlen := 0
	hashArr := [26]int {}
	maxF := 0

	i, j := 0, 0
	for j < len(s) {
		hashArr[s[j]-'A']++
		if maxF < hashArr[s[j]-'A'] {
			maxF = hashArr[s[j]-'A']
		}

		if (j-i+1) - maxF > k {
			hashArr[s[i]-'A']--
			i++
		}

		if (j-i+1)-maxF <= k {
			if maxlen < j-i+1 {
				maxlen = j-i+1
			}
		}
		j++
	}
	return maxlen
}