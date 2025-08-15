package main

import (
	"fmt"
	"math"
)

func main() {
	a := 123456

	fmt.Printf("number of digits in %d using the brute force approach: %d\n", a, brute(a))
	fmt.Printf("number of digits in %d using the log function: %d\n", a, mathLog(a))
}

func brute(a int) int {
	count := 0
	for a > 0 {
		count++
		a = a / 10
	}
	return count
}

func mathLog(a int) int {
	count := int(math.Log10(float64(a)))+1
	return count
}