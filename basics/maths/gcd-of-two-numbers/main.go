package main

import "fmt"

func main() {
	num1 := 20
	num2 := 15

	fmt.Printf("gcd of %d and %d is: %d\n", num1, num2, gcd(num1, num2))
	fmt.Printf("gcd of %d and %d using Euclidean algorithm is: %d\n", num1, num2, euclidean(num1, num2))
}

func gcd(a, b int) int {
	for i := min(a, b); i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}


func euclidean(a, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a == 0 {
		return b
	}
	return a
}