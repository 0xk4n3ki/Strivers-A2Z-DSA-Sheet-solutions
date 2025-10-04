package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {1, 2, 3, 4, 5, 6, 1}
	k := 3
	fmt.Printf("ans: %d\n", brute(arr, k))
	fmt.Printf("ans: %d\n", optimal(arr, k))
}

func brute(arr []int, k int) int {
	ans := math.MinInt
	n := len(arr)

	for i := 0; i <= k; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += arr[j]
		}
		for j := n-1; j >= n-k+i; j-- {
			sum += arr[j]
		}

		if sum > ans {
			ans = sum
		}
	}
	return ans
}

func optimal(arr []int, k int) int {
	ans := 0
	for i := 0; i < k; i++ {
		ans += arr[i]
	}

	sum := ans
	n := len(arr)
	for i := 0; i < k; i++ {
		sum -= arr[k-i-1]
		sum += arr[n-1-i]
		
		if sum > ans {
			ans = sum
		}
	}
	return ans
}