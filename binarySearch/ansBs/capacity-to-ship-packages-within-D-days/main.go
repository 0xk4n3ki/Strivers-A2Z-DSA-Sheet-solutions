package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {5, 4, 5, 2, 3, 4, 5, 6}
	days := 5

	bs(arr, days)

	arr = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days = 1

	bs(arr, days)
}

func bs(arr []int, limit int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}

	low, high := maxEle(arr), sum

	for low <= high {
		mid := (low + high)/2

		days := calc(arr, mid)
		fmt.Printf("mid: %d, days: %d\n", mid, days)

		if days <= limit {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("arr: %v, days limit: %d, weight capacity: %d\n", arr, limit, low)
}

func calc(arr []int, limit int) int {
	days := 1
	sum := 0
	for i := 0; i < len(arr); i++ {
		if sum + arr[i] <= limit {
			sum += arr[i]
		}else {
			days++
			sum = arr[i]
		}
	}
	return days
}

func maxEle(arr []int) int {
	max := math.MinInt

	for _, val := range arr {
		if max < val {
			max = val
		}
	}
	return max
}