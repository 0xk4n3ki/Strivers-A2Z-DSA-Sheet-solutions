package main

import (
	"fmt"
)

func main() {
	s := "aababbcaacc"
	k := 2
	fmt.Printf("ans: %d\n", brute(s, k))
	fmt.Printf("ans: %d\n", longest(s, k))

	s = "abcddefg"
	k = 3
	fmt.Printf("ans: %d\n", brute(s, k))
	fmt.Printf("ans: %d\n", longest(s, k))
}


func longest(s string, k int) int {
	ans := 0
	left, right := 0, 0
	n := len(s)
	charMap := map[byte]int {}

	for right < n {
		charMap[s[right]]++
		
		for len(charMap) > k {
			charMap[s[left]]--
			if charMap[s[left]] == 0 {
				delete(charMap, s[left])
			}
			left++
		}

		if right - left + 1 > ans {
			ans = right - left + 1
		}
		right++
	}
	return ans
}



func brute(s string, k int) int {
	max := 0

	for i := 0; i < len(s); i++ {
		charMap := map[byte]int {}
		for j := i; j < len(s); j++ {
			charMap[s[j]]++

			if len(charMap) > k {
				break
			}

			if j-i+1 > max {
				max = j-i+1
			}
		}
	}
	return max
}