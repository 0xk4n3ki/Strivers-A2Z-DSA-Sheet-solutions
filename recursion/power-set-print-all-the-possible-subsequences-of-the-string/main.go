package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abc"
	ans := subSeq(s)
	fmt.Printf("ans: %v, len: %d\n", ans, len(ans))

	sol := bitMani(s)
	fmt.Printf("ans: %v, len: %d\n", sol, len(sol))
}

// backtracking
func subSeq(s string) []string {
	ans := []string {}
	// substr := []

	solve(0, s, "", &ans)
	sort.Strings(ans)

	return ans
}

func solve(index int, s string, substr string, ans *[]string) {
	if index == len(s) {
		*ans = append(*ans, substr)
		return
	}

	solve(index+1, s, substr+string(s[index]), ans)
	solve(index+1, s, substr, ans)
}



// bit manipulation
func bitMani(s string) []string {
	ans := []string {}
	n := len(s)
	for i := 0; i < 1 << n; i++ {
		subStr := []rune {}

		for j := 0; j < n; j++ {
			if i & (1 << j) != 0 {
				subStr = append(subStr, rune(s[j]))
			}
		}
		ans = append(ans, string(subStr))
	}

	return ans
}