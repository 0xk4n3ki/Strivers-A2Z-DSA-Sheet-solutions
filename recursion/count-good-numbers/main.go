package main

import (
	"fmt"
	// "math"
)

func main() {
	fmt.Printf("n: 4, ans: %d\n", countGoodNumbers(4))
	fmt.Printf("n: 1924, ans: %d\n", countGoodNumbers(1924))
	fmt.Printf("n: 50, ans: %d\n", countGoodNumbers(50))
}


const MOD = 1_000_000_007
func countGoodNumbers(n int64) int {
    if n == 1 {
		return 5
	}

	even := (n+1)/2
	odd := n/2

	res := (myPow(5, even) * myPow(4, odd)) % MOD
	return int(res)
}

func myPow(x, n int64) int64 {
	if n == 0 {
		return 1
	}

	x = x % MOD

	if n % 2 == 0 {
		half := myPow(x, n/2) % MOD
		return (half*half) % MOD
	}else {
		return (x * myPow(x, n-1)) % MOD
	}
}