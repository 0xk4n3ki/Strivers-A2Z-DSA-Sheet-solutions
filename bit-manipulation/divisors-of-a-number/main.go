package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("n: %d, ans: %v\n", 6, getPrime(6))
	fmt.Printf("n: %d, ans: %v\n", 8, getPrime(8))
}

func getPrime(n int) []int {
	ans := []int {}

	for i := 1; i*i <= n; i++ {
		if n % i == 0 {
			ans = append(ans, i)

			if n/i != i {
				ans = append(ans, n/i)
			}
		}
	}

	sort.Ints(ans)
	return ans
}