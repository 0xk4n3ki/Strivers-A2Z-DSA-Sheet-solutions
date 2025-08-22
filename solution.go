package main

import (
	"fmt"
)

// [OLD] https://takeuforward.org/data-structure/find-intersection-of-two-linked-lists/
// [OLD] https://takeuforward.org/data-structure/merge-overlapping-sub-intervals/
// [OLD] https://takeuforward.org/binary-search/find-peak-element-ii
// [NEW] https://takeuforward.org/data-structure/aggressive-cows-detailed-solution/
// [NEW] https://leetcode.com/problems/roman-to-integer/description/
// [NEW] https://takeuforward.org/data-structure/combination-sum-ii-find-all-unique-combinations/
// [NEW] https://takeuforward.org/data-structure/palindrome-partitioning/
// [NEW] https://takeuforward.org/plus/dsa/problems/xor-of-numbers-in-a-given-range
// [NEW] https://takeuforward.org/data-structure/check-for-balanced-parentheses/
// [NEW] https://leetcode.com/problems/online-stock-span/description/
// [NEW] https://leetcode.com/problems/longest-repeating-character-replacement/description/

func main() {
	s := "ABAA"
	k := 0
	ans := characterReplacement(s, k)
	fmt.Printf("ans: %d\n", ans)

	s = "AABABBA"
	k = 1
	ans = characterReplacement(s, k)
	fmt.Printf("ans: %d\n", ans)
}

func characterReplacement(s string, k int) int {
	maxlen := 0
	i, j := 0, 0
	maxf := 0
	charmap := [26]int {}

	for j < len(s) {
		charmap[s[j]-'A']++

		if maxf < charmap[s[j]-'A'] {
			maxf = charmap[s[j]-'A']
		}

		n := j-i+1-maxf

		if n > k {
			charmap[s[i]-'A']--
			i++
		}

		if j-i+1-maxf <= k {
			if j-i+1 > maxlen {
				maxlen = j-i+1
			}
		}
		
		j++
	}
	return maxlen
}