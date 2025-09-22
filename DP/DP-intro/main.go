package main

import "fmt"

func main() {
	n := 10
	dp := make([]int, n+1)
	fmt.Printf("ans: %d\n", memoization(n, dp))
	fmt.Printf("dp: %v\n", dp)

	dp = make([]int, n+1)
	fmt.Printf("ans: %d\n", tabulation(n, dp))
	fmt.Printf("dp: %v\n", dp)

	fmt.Printf("ans: %d\n", spaceOpt(n))
}

func spaceOpt(n int) int {
	prev2, prev := 0, 1

	for i := 2; i <= n; i++ {
		curr := prev + prev2
		prev2 = prev
		prev = curr
	}
	return prev
}

func tabulation(n int, dp []int) int {
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}

func memoization(n int, dp []int) int {
	if n <= 1 {
		dp[n] = n
		return n
	}

	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = memoization(n-1, dp) + memoization(n-2, dp)
	return dp[n]
}

