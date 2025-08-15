package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {5, 5, 5, 5}
	k := 2

	bs(arr, k)

	arr = []int {10, 20, 30, 40}
	k = 2
	bs(arr, k)
}

func bs(arr []int, k int) {
	low, high := calcLimits(arr)

	for low <= high {
		mid := (low + high)/2

		painters := calcNum(arr, mid)

		if painters > k {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, k: %d, time: %d\n", arr, k, low)
}

func calcNum(arr []int, mid int) int {
	count := 1
	sum := 0

	for i := 0; i < len(arr); i++ {
		if sum + arr[i] > mid {
			count++
			sum = arr[i]
		}else {
			sum += arr[i]
		}
	}
	return count
}

func calcLimits(arr []int) (int, int) {
	sum, maxEle := 0, math.MinInt

	for _, val := range arr {
		sum += val
		if val > maxEle {
			maxEle = val
		}
	}
	return maxEle, sum
}