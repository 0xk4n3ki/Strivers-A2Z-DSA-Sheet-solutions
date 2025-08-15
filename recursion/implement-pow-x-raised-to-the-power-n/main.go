package main

import (
	"fmt"
	"math"
)

func main() {
	n := myPow(2, 10)
	fmt.Printf("2^10: %f\n", n)

	n = myPow(2.1, 3)
	fmt.Printf("2.1^3: %f\n", n)
}

// func myPow(x float32, n int) float32 {
// 	if n == 0 {
// 		return 1
// 	}

// 	return x * myPow(x, n-1)
// }

func myPow(x float32, n int) float32 {
	if n < 0 {
		// special case of math.MinInt to avoid overflow
		if n == math.MinInt {
			return float32(1) / (x * myPow(x, math.MaxInt))
		}
		return float32(1) / myPow(x, -n)
	}

	// base case
	if n == 0 {
		return 1
	}

	// if n is even: x^n = (x*x)^(n/2)
	if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	}

	// if n is odd: x^n = x * x^(n-1)
	return x * myPow(x, n-1)
}
