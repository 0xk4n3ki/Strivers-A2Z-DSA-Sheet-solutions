package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "this is an amazing program"
	s1 := "This is decent"

	reverse(s)
	reverse(s1)

	usingPointers(s)
	usingPointers(s1)
}

func reverse(s string) {
	a := strings.Split(s, " ")

	fmt.Printf("string: %s, ", s)
	var builder strings.Builder

	for i := len(a)-1; i > 0; i-- {
		builder.WriteString(a[i])
		builder.WriteRune(' ')
	}
	builder.WriteString(a[0])
	fmt.Printf("reversed string: %s\n", builder.String())
}

func usingPointers(s string) {
	left, right := 0, len(s)-1
	tmp, ans := "", ""

	// append word in front of answer
	for left <= right {
		ch := s[left]

		if ch != ' ' {
			tmp += string(ch)
		}else {
			if ans != "" {
				ans = tmp + " " + ans
			}else {
				ans = tmp
			}
			tmp = ""
		}
		left++
	}

	// if something in tmp, add that too to ans
	if tmp != "" {
		if ans != "" {
			ans = tmp + " " + ans
		}else {
			ans = tmp
		}
	}

	fmt.Printf("s: %s, ans: %s\n", s, ans)
}