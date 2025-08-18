package main

import (
	"fmt"
)

// [OLD] https://leetcode.com/problems/power-of-two/description/
// [OLD] https://takeuforward.org/data-structure/rotate-image-by-90-degree/
// [OLD] https://takeuforward.org/data-structure/merge-overlapping-sub-intervals/
// [NEW] https://takeuforward.org/arrays/painters-partition-problem/
// [NEW] https://takeuforward.org/data-structure/check-if-two-strings-are-anagrams-of-each-other/
// [NEW] https://takeuforward.org/data-structure/power-set-print-all-the-possible-subsequences-of-the-string/
// [NEW] https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/
// [NEW] https://leetcode.com/problems/minimum-bit-flips-to-convert-number/description/
// [NEW] postfix-to-infix: ab-de+f*/
// [NEW] https://takeuforward.org/data-structure/sliding-window-maximum/
// [NEW] https://takeuforward.org/data-structure/length-of-longest-substring-without-any-repeating-character/

func main() {
	s := "abcabcbb"
	ans := long(s)
	fmt.Printf("ans: %d\n", ans)

	s = "bbbbb"
	ans = long(s)

	fmt.Printf("ans: %d\n", ans)
}

func long(s string) int {
	maxLen := 0
	i, j := 0, 0

	charMap := map[byte]int {}
	
	for j < len(s) {

		if index, ok := charMap[s[j]]; ok {
			if index+1 > i {
				i = index+1
			}
		}


		l := j-i+1
		if l > maxLen {
			maxLen = l
		}
		charMap[s[j]] = j
		j++
	}
	return maxLen
}