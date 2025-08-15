package main

import (
	"fmt"
	"math"
)

func main() {
	roses := 8
	arrDays := []int {7, 7, 7, 7, 13, 11, 12, 7}
	bouq := 2
	rosesToBouq := 3

	bruteforce(roses, arrDays, bouq, rosesToBouq)
	binarySearch(arrDays, bouq, rosesToBouq)

	roses = 5
	arrDays = []int {1, 10, 3, 10, 2}
	bouq = 3
	rosesToBouq = 2
	bruteforce(roses, arrDays, bouq, rosesToBouq)
	binarySearch(arrDays, bouq, rosesToBouq)
}

func binarySearch(arr []int, bouq, rosesToBouq int) {
	n := len(arr)
	if rosesToBouq * bouq > n {
		fmt.Printf("arr: %v, num of roses; %d, bouq req: %d, roses per bouq: %d, ans: %d\n", arr, n, bouq, rosesToBouq, -1)
		return
	}

	high, low := maxArr(arr), minArr(arr)
	for low <= high {
		mid := (low+high)/2

		if possible(arr, mid, bouq, rosesToBouq) {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("arr: %v, num of roses; %d, bouq req: %d, roses per bouq: %d, ans: %d\n", arr, n, bouq, rosesToBouq, low)
}

func bruteforce(n int, arr []int, bouq, rosesToBouq int) {
	if n < bouq * rosesToBouq {
		fmt.Printf("arr: %v, num of roses; %d, bouq req: %d, roses per bouq: %d, ans: %d\n", arr, n, bouq, rosesToBouq, -1)
		return
	}

	minDay := minArr(arr)
	maxDay := maxArr(arr)
	// fmt.Printf("minDay: %d, maxDay: %d\n", minDay, maxDay)

	var i int
	for i = minDay; i <= maxDay; i++ {
		if possible(arr, i, bouq, rosesToBouq) {
			break
		}
	}
	fmt.Printf("arr: %v, num of roses; %d, bouq req: %d, roses per bouq: %d, ans: %d\n", arr, n, bouq, rosesToBouq, i)
}

func possible(arr[] int, day, bouq, rosesToBouq int) bool {
	n := len(arr)
	bouqCount := 0
	count := 0
	for j := 0; j < n; j++{
		if arr[j] <= day {
			count++
		}else {
			bouqCount += count/rosesToBouq
			count = 0
		}
		if bouqCount >= bouq {
			break
		}
		// fmt.Printf("i: %d, arr[%d]: %d, count: %d, bouqCount: %d\n", day, j, arr[j], count, bouqCount)
	}
	bouqCount += count/rosesToBouq

	return bouqCount >= bouq
}

func minArr(arr []int) int {
	min := math.MaxInt

	for _, val := range arr {
		if min > val {
			min = val
		}
	}
	return min
}

func maxArr(arr []int) int {
	max := math.MinInt

	for _, val := range arr {
		if max < val {
			max = val
		}
	}
	return max
}