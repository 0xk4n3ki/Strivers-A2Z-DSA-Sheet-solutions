package main

import (
	"fmt"
	"math"
)

func main() {
	s := "42"
	atoi(s)

	s = " -042"
	atoi(s)

	s = "1337c0d3"
	atoi(s)

	s = "+-12"
	atoi(s)

	s = "-91283472332"
	atoi(s)
}


func atoi(s string) {
	n := len(s)

	if n == 0 {
		fmt.Printf("s: %s, ans: %d\n", s, 0)
		return
	}

	max, min := math.MaxInt32, math.MinInt32

	i := 0

	// skip leading whitespaces
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		fmt.Printf("s: %s, ans: %d\n", s, 0)
		return
	}

	sign := 1
	if s[i] == '+' {
		i++
	}else if s[i] == '-' {
		sign = -1
		i++
	}

	res := 0
	for i < n && 0 <= s[i] - '0' && s[i] - '0' <= 9 {
		digit := s[i] - '0'
		res = res * 10 + int(digit)

		if sign * res <= min {
			fmt.Printf("s: %s, ans: %d\n", s, min)
			return
		}
		if sign * res >= max {
			fmt.Printf("s: %s, ans: %d\n", s, max)
			return
		}

		i++
	}

	fmt.Printf("s: %s, ans: %d\n", s, sign*res)
}