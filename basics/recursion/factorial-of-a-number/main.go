package main

import "fmt"

func main() {
	n1 := 5
	n2 := 6

	fmt.Printf("factorial of %d: %d\n", n1, factorial(n1))
	fmt.Printf("factorial of %d: %d\n", n2, factorial(n2))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}