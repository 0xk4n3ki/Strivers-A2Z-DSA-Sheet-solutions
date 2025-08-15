package main

import "fmt"

func main() {
	n1 := 5
	n2 := 6

	fmt.Printf("fibonacci series upto %d: %v\n", n1, fibonacci(n1))
	fmt.Printf("fibonacci series upto %d: %v\n", n2, fibonacci(n2))
}

func fibonacci(n int) []int {
	if n == 0 {
		return []int {0}
	}else if n == 1 {
		return []int {0, 1}
	}
	result := fibonacci(n-1)
	next := result[len(result)-1] + result[len(result)-2]
	return append(result, next)
}