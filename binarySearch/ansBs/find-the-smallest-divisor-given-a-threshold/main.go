package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {1, 2, 3, 4, 5}
	limit := 8

	bs(arr, limit)

	arr = []int {8, 4, 2, 3}
	limit = 10

	bs(arr, limit)
}

func bs(arr []int, limit int) {
	low, high := 1, max(arr)

	for low <= high {
		mid := (low + high)/2
		
		sum := calc(arr, mid)
		// fmt.Printf("mid: %d, sum: %d\n", mid, sum)

		if sum > limit {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, limit: %d, divisor: %d\n", arr, limit, low)
}

func calc(arr []int, mid int) int {
	sum := 0
	for _, val := range arr {
		sum += ceil(val, mid)
	}
	return sum
}

func ceil(val, mid int) int {
	res := val/mid
	if val % mid != 0 {
		res += 1
	}
	return res
}

func max(arr []int) int {
	res := math.MinInt

	for _, val := range arr {
		if val > res {
			res = val
		}
	}
	return res
}