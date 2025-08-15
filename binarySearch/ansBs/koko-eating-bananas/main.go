package main

import (
	"fmt"
	"math"
)

func main() {
	n, h := 4, 8
	arr := []int {7, 15, 6, 3}
	bs(n, h, arr)

	n, h = 5, 5
	arr = []int {25, 12, 8, 14, 19}
	bs(n, h, arr)
}

func bs(n, h int, arr []int) {
	low, high := 1, max(arr)
	ans := -1

	for low <= high {
		mid := (low+high)/2

		time := 0
		for i := 0; i < len(arr); i++ {
			time += ceil(arr[i], mid)
		}

		if time <= h {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	ans = low

	fmt.Printf("arr: %v, n: %d, h: %d, ans: %d\n", arr, n, h, ans)
}

func ceil(a, b int) int {
	sum := 0

	sum += a/b
	if a % b != 0 {
		sum += 1
	}
	return sum
}

func max(arr []int) int {
	maxNum := math.MinInt
	for _, val := range arr {
		if val > maxNum {
			maxNum = val
		}
	}
	return maxNum
}