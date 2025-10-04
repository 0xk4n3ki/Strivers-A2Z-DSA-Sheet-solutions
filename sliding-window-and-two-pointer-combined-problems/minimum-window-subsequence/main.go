package main

import (
	"fmt"
	"math"
)

func main() {
	s, t := "geeksforgeeks", "eksrg"
	fmt.Printf("ans: %s\n", brute(s, t))
	fmt.Printf("ans: %s\n", twoPointer(s, t))
	// fmt.Printf("ans: %s\n", minWindow(s, t))

	s, t = "abcdebdde", "bde"
	fmt.Printf("ans: %s\n", brute(s, t))
	fmt.Printf("ans: %s\n", twoPointer(s, t))
	// fmt.Printf("ans: %s\n", minWindow(s, t))

	s, t = "ad", "b"
	fmt.Printf("ans: %s\n", brute(s, t))
	fmt.Printf("ans: %s\n", twoPointer(s, t))
	// fmt.Printf("ans: %s\n", minWindow(s, t))
}

// func minWindow(s, t string) string {
// 	n, m := len(s), len(t)
	
// }

func twoPointer(s, t string) string {
	n, m := len(s), len(t)
	min, start := math.MaxInt, -1

	for i := 0; i < n; i++ {
		if s[i] == t[0] {
			sp, tp := i, 0

			for sp < n && tp < m {
				if s[sp] == t[tp] {
					tp++
				}
				sp++
			}

			if tp == m {
				// end := sp-1
				// tp = m-1

				// for end >= i {
				// 	if s[end] == t[tp] {
				// 		tp--
				// 	}
				// 	if tp < 0 {
				// 		break
				// 	}
				// 	end--
				// }

				// len := sp-end
				// if len < min {
				// 	min = len
				// 	start = end
				// }
				len := sp - i
				if len < min {
					min = len
					start = i
				}
			}
		}
	}
	if start == -1 {
		return ""
	}
	return s[start:start+min]
}


func brute(s, t string) string {
	n := len(s)
	min := math.MaxInt
	start := -1

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isSub(s[i:j+1], t) {
				if j-i+1 < min {
					min = j-i+1
					start = i
				}
				break
			}
		}
	}
	if start == -1 {
		return ""
	}
	return s[start:start+min]
}

func isSub(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return j == len(t)
}