package main

import (
	"fmt"
	"math"
)

func main() {
	s := "2345"
	n := myAtoi(s)
	fmt.Printf("s: %s, n: %d\n", s, n)
}

func myAtoi(s string) int {
	n := len(s)
	if n == 1 {
		return int(s[0] - '0')
	}	

	front := myAtoi(s[1:])

	tens := math.Pow10(n-1)

	res := int(s[0]-'0') * int(tens) + front

	return res
}