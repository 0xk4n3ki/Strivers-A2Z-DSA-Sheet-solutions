package main

import (
	"fmt"
	"math"
)

func main() {
	energy := []int {30, 10, 60, 10, 60, 50}

	dp := make([]int, len(energy))
	for i := 0; i < len(energy); i++ {
		dp[i] = -1
	}
	fmt.Printf("ans: %d\n", memoization(len(energy)-1, energy, dp))
	fmt.Printf("ans: %d\n", tabulation(energy))
	fmt.Printf("ans: %d\n", spaceOpt(energy))
}

func spaceOpt(energy []int) int {
	n := len(energy)
	prev, prev2 := 0, 0
	for i := 1; i < n; i++ {
		jumpTwo := math.MaxInt
		jumpOne := prev + int(math.Abs(float64(energy[i])-float64(energy[i-1])))
		if i > 1 {
			jumpTwo = prev2 + int(math.Abs(float64(energy[i])-float64(energy[i-2])))
		}

		
		curr := min(jumpOne, jumpTwo)
		fmt.Printf("i: %d, prev: %d, prev2: %d, curr: %d\n", i, prev, prev2, curr)
		prev2 = prev
		prev = curr
	}
	return prev
}

func tabulation(energy []int) int {
	n := len(energy)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < n; i++ {
		jumpTwo := math.MaxInt
		jumpOne := dp[i-1] + int(math.Abs(float64(energy[i])-float64(energy[i-1])))
		if i > 1 {
			jumpTwo = dp[i-2] + int(math.Abs(float64(energy[i])-float64(energy[i-2])))
		}
		dp[i] = min(jumpOne, jumpTwo)
	}
	return dp[n-1]
}

func memoization(ind int, energy, dp []int) int {
	if ind == 0 {
		return 0
	}
	if dp[ind] != -1 {
		return dp[ind]
	}

	jumpTwo := math.MaxInt
	jumpOne := memoization(ind-1, energy, dp) + int(math.Abs(float64(energy[ind])-float64(energy[ind-1])))

	if ind > 1 {
		jumpTwo = memoization(ind-2, energy, dp) + int(math.Abs(float64(energy[ind])-float64(energy[ind-2])))
	}

	dp[ind] = min(jumpOne, jumpTwo)
	return dp[ind]
}