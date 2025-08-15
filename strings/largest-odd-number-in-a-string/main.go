package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("ans: %s\n", largest("5347"))
	fmt.Printf("ans: %s\n", largest("0214638"))

	fmt.Printf("ans: %s\n", second("5347"))
	fmt.Printf("ans: %s\n", second("0214638"))
}

func second(s string) string{
	n := toInt(s)

	for n > 0 {
		if n % 2 != 0 {
			break
		}else {
			n = n / 10
		}
	}
	// fmt.Printf("s: %s, ans: %d\n", s, n)

	n = rev(n)

	var builder strings.Builder

	for n > 0 {
		rem := n % 10
		builder.WriteRune(rune('0' + rem))
		n /= 10
	}
	return builder.String()
}

func rev(n int) int {
	res := 0

	for n > 0 {
		rem := n % 10
		res = (res*10) + rem
		n /= 10
	}
	return res
}

func toInt(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		res = (res*10) + int(s[i]-'0')
	}

	return res
}

func largest(s string) string {

	var i int
	for i = len(s)-1; i >= 0; i-- {
		// n := int(s[i])
		if (s[i]-'0')%2 != 0 {
			
			start := 0

			for start < i && s[start] == '0' {
				start++
			}

			return s[start: i+1]
		}
	}
	return ""
}