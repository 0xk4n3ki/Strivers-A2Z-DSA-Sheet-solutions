package main

import (
	"fmt"
	"math"
)

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Printf("ans: %s\n", brute(s, t))
	fmt.Printf("ans: %s\n", optimal(s, t))

	s, t = "a", "a"
	fmt.Printf("ans: %s\n", brute(s, t))
	fmt.Printf("ans: %s\n", optimal(s, t))

	s, t = "a", "aa"
	fmt.Printf("ans: %s\n", brute(s, t))
	fmt.Printf("ans: %s\n", optimal(s, t))
}

func optimal(s, t string) string {
	min, start := math.MaxInt, -1
	cnt := 0
	hashMap := map[byte]int {}
	left, right := 0, 0

	for i := 0; i < len(t); i++ {
		hashMap[t[i]]++
	}

	for right < len(s) {
		if hashMap[s[right]] > 0 {
			cnt++
		}
		hashMap[s[right]]--

		for cnt == len(t) {
			if right-left+1 < min {
				min = right-left+1
				start = left
			}
			hashMap[s[left]]++
			if hashMap[s[left]] > 0 {
				cnt--
			}
			left++
		}

		right++
	}
	if start == -1 {
		return ""
	}
	return s[start:start+min]
}

func brute(s, t string) string {
	min, start := math.MaxInt, -1
	
	for i := 0; i < len(s); i++ {
		hash := [256]int {}
		for j := 0; j < len(t); j++ {
			hash[t[j]]++
		}

		cnt := 0
		for j := i; j < len(s); j++ {
			if hash[s[j]] > 0 {
				cnt++
			}
			hash[s[j]]--
			if cnt == len(t) {
				if j-i+1 < min {
					min = j-i+1
					start = i
					break
				}
			}
		}
	}
	// fmt.Printf("%d, %d\n", start, min)
	if start == -1 {
		return ""
	}
	return s[start:start+min]
}