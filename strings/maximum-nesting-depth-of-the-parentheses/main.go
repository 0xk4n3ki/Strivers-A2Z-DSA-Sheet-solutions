package main

import (
	"fmt"
	"math"
)

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	try(s)

	s = "(1)+((2))+(((3)))"
	try(s)

	s = "()(())((()()))"
	try(s)
}

func try(s string) {
	count := 0

	num := 0 
	for _, c := range s {
		if c == '(' {
			num++
			if count < num {
				count = num
			}
		}else if c == ')' {
			num--
		}
	}
	fmt.Printf("s: %s, ans: %d\n", s, count)
}