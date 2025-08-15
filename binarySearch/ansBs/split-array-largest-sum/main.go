package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {1, 2, 3, 4, 5}
	k := 3

	bs(arr, k)

	arr = []int {3, 5, 1}
	k = 3
	bs(arr, k)
}

func bs(arr []int, k int) {
	low, high := findLimits(arr)

	for low <= high {
		mid := (low + high)/2

		count := calcSubArr(arr, mid)

		if count > k {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, k: %d, num of subarrays: %d\n", arr, k, low)
}

func calcSubArr(arr []int, mid int) int {
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

func findLimits(arr []int) (int, int) {
	maxEle, sum := math.MinInt, 0

	for _, val := range arr {
		sum += val

		if maxEle < val {
			maxEle = val
		}
	}
	return maxEle, sum
}