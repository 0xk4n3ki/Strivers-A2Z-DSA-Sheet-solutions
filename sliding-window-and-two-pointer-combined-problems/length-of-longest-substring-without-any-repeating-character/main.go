package main

import "fmt"

func main() {
	s := "abcabcbb"
	ans := longSub(s)
	fmt.Printf("ans: %d\n", ans)

	s = "abcaabcdba"
	ans = longSub(s)
	fmt.Printf("ans: %d\n", ans)

	s = "pwwkew"
	ans = longSub(s)
	fmt.Printf("ans: %d\n", ans)

	s = " "
	ans = longSub(s)
	fmt.Printf("ans: %d\n", ans)

	s = "abba"
	ans = longSub(s)
	fmt.Printf("ans: %d\n", ans)
}

func longSub(s string) int {
	n := len(s)
	ans := 0
	i, j := 0, 0
	hashSet := map[byte]int{}

	for j < n {
		if _, ok := hashSet[s[j]]; ok {
			i = max(hashSet[s[j]]+1, i)
		}
		hashSet[s[j]] = j
		ans = max(ans, j-i+1)
		j++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
