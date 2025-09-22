package main

import (
	"fmt"
	"math"
)

func main() {
	energy := []int {30, 10, 60, 10, 60, 50}
	k := 2

	n := len(energy)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	fmt.Printf("ans: %d\n", memoization(n-1, energy, k, dp))

	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	fmt.Printf("ans: %d\n", tabular(n, energy, k, dp))
}

func tabular(n int, energy []int, k int, dp []int) int{
	dp[0] = 0
	for i := 1; i < n; i++ {
		minSteps := math.MaxInt
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				jump := dp[i-j] + int(math.Abs(float64(energy[i] - energy[i-j])))
				minSteps = min(jump, minSteps)
			}
		}
		dp[i] = minSteps
	}
	return dp[n-1]
}

func memoization(ind int, energy []int, k int, dp []int) int {
	if ind == 0 {
		return 0
	}
	if dp[ind] != -1 {
		return dp[ind]
	}

	minSteps := math.MaxInt
	for j := 1; j <= k; j++ {
		if ind - j >= 0 {
			jump := memoization(ind-j, energy, k, dp) + int(math.Abs(float64(energy[ind] - energy[ind-j])))
			minSteps = min(jump, minSteps)
		}
	}
	dp[ind] = minSteps
	return dp[ind]
}