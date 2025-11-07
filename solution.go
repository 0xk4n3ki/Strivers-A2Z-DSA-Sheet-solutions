package main

import (
	"fmt"
)

// [OLD] https://leetcode.com/problems/largest-odd-number-in-string/description/
// [OLD] https://takeuforward.org/data-structure/quick-sort-algorithm/
// [OLD] https://takeuforward.org/data-structure/segregate-even-and-odd-nodes-in-linkedlist
// [NEW] https://takeuforward.org/arrays/painters-partition-problem/
// [NEW] https://leetcode.com/problems/roman-to-integer/description/
// [NEW] https://takeuforward.org/data-structure/combination-sum-ii-find-all-unique-combinations/
// [NEW] https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/
// [NEW] https://leetcode.com/problems/count-primes/description/
// [NEW] https://takeuforward.org/plus/dsa/problems/postfix-to-prefix-conversion
// [NEW] https://leetcode.com/problems/sum-of-subarray-minimums/description/
// [NEW] https://leetcode.com/problems/minimum-window-substring/


func main() {
	s := "III"
	fmt.Printf("s: %s, ans: %d\n", s, calc(s))

	s = "LVIII"
	fmt.Printf("s: %s, ans: %d\n", s, calc(s))

	s = "MCMXCIV"
	fmt.Printf("s: %s, ans: %d\n", s, calc(s))
}

func calc(s string) int {
	values := map[byte]int {
		'I' : 1,
		'V' : 5,
		'X': 10,
		'L' : 50,
		'C' : 100,
		'D' : 500,
		'M' : 1000,
	}

	ans := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[s[i]] < values[s[i+1]] {
			ans -= values[s[i]]
		}else {
			ans += values[s[i]]
		}
	}
	return ans
}


// func calc(s string) int {
// 	ans := 0

// 	for i := 0; i < len(s); i ++ {
// 		c := s[i]

// 		switch c {
// 		case 'I':
// 			if i+1 < len(s) && s[i+1] == 'V' {
// 				ans += 4
// 				i++
// 			}else if i+1 < len(s) && s[i+1] == 'X' {
// 				ans += 9
// 				i++
// 			}else {
// 				ans += 1
// 			}
// 		case 'V':
// 			ans += 5
// 		case 'X':
// 			if i+1 < len(s) && s[i+1] == 'L' {
// 				ans += 40
// 				i++
// 			}else if i+1 < len(s) && s[i+1] == 'C' {
// 				ans += 90
// 				i++
// 			}else {
// 				ans += 10
// 			}
// 		case 'L':
// 			ans += 50
// 		case 'C':
// 			if i+1 < len(s) && s[i+1] == 'D' {
// 				ans += 400
// 				i++
// 			}else if i+1 < len(s) && s[i+1] == 'M' {
// 				ans += 900
// 				i++
// 			}else {
// 				ans += 100
// 			}
// 		case 'D':
// 			ans += 500
// 		case 'M':
// 			ans += 1000
// 		}
// 	}
// 	return ans
// }