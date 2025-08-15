package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "((()))"
	b := "()(()())(())"
	c := "(()())(())"

	// remove(a)
	// remove(b)

	fmt.Printf("%s\n", removeOuterParentheses(a))
	fmt.Printf("%s\n", removeOuterParentheses(b))
	fmt.Printf("%s\n", removeOuterParentheses(c))

	fmt.Printf("%s\n", usingBuilder(a))
	fmt.Printf("%s\n", usingBuilder(b))
	fmt.Printf("%s\n", usingBuilder(c))
}

// func remove(a string) {
// 	count := 0

// 	for _, i := range a {
// 		if count == 0 && i == '(' {
// 			count++
// 			continue
// 		}else if count == 1 && i == ')' {
// 			count--
// 			continue
// 		}else if i == '(' {
// 			count++
// 		}else if i == ')' {
// 			count--
// 		}
// 		fmt.Printf("%c", i)
// 	}
// 	fmt.Println()
// }

func removeOuterParentheses(s string) string {
    count := 0

    a := ""

	for _, i := range s {
		if i == '(' {
			if count > 0 {
				a += string(i)
			}
			count++
		}else if i == ')' {
			count--
			if count > 0 {
				a += string(i)
			}
		}
	}
	return a
}

func usingBuilder(s string) string {
	var builder strings.Builder

	count := 0

	for _, c := range s {
		if c == '(' {
			if count > 0 {
				builder.WriteRune(c)
			}
			count++
		}else if c == ')' {
			count--
			if count > 0 {
				builder.WriteRune(c)
			}
		}
	}
	return builder.String()
}