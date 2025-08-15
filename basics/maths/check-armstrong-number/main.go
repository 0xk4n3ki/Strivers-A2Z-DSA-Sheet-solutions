package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := 153
	num2 := 371
	num3 := 372

	fmt.Printf("%d is a armstrong number: %v\n", num1, check(num1))
	fmt.Printf("%d is a armstrong number: %v\n", num2, check(num2))
	fmt.Printf("%d is a armstrong number: %v\n", num3, check(num3))
}

func check(a int) bool {
	dup := a
	sum := 0

	if a == 0 {
		return true
	}

	len := int(math.Log10(float64(a)))+1

	for a > 0 {
		base := a % 10
		sum += pow(base, len)
		a = a / 10
	}
	return dup == sum
}

func pow(base, p int) int {
	result := 1
	for p > 0 {
		result *= base
		p--
	}
	return result
}