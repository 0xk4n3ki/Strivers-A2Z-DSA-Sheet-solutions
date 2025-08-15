package main

import (
	"fmt"
	"strings"
)

func main() {
	s, goal := "abcde", "cdeab"
	fmt.Printf("s: %s, goal: %s, ans: %t\n", s, goal, rotateString(s, goal))
	fmt.Printf("s: %s, goal: %s, ans: %t\n", s, goal, conCheck(s, goal))
	fmt.Printf("s: %s, goal: %s, ans: %t\n", s, goal, kmp(s, goal))

	s, goal = "abcde", "abced"
	fmt.Printf("s: %s, goal: %s, ans: %t\n", s, goal, rotateString(s, goal))
	fmt.Printf("s: %s, goal: %s, ans: %t\n", s, goal, conCheck(s, goal))
	fmt.Printf("s: %s, goal: %s, ans: %t\n", s, goal, kmp(s, goal))
}


// concatenation check
func conCheck(s, goal string) bool {
	n1, n2 := len(s), len(goal)
	if n1 != n2 {
		return false
	}
	
	conStr := s + s

	return strings.Contains(conStr, goal)
}


// bruteforce - trying all possible rotations
func rotateString(s, goal string) bool {
	n1, n2 := len(s), len(goal)

	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		newStr := rotate(s, i)
		// fmt.Printf("newstr: %s\n", newStr)

		if newStr == goal {
			return true
		}
	}
	return false
}

func rotate(s string, i int) string {
	res := ""

	for j := i; j < len(s); j++ {
		res += string(s[j])
	}

	for j := 0; j < i; j++ {
		res += string(s[j])
	}
	return res
}