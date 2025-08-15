package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 453979354
	b := 1234

	fmt.Printf("%d is palindrome: %v\n", a, check(a))
	fmt.Printf("%d is palindrome: %v\n", b, check(b))

	fmt.Printf("%d is palindrome using check2 func: %v\n", a, check2(a))
	fmt.Printf("%d is palindrome using check2 func: %v\n", b, check2(b))
}

func check(a int) bool {
	s := strconv.Itoa(a)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func check2(a int) bool {
	dup := a
	rev := 0
	for dup > 0 {
		ld := dup % 10
		rev = rev * 10 + ld
		dup = dup / 10
	}
	if rev == a {
		return true
	}else {
		return false
	}
}