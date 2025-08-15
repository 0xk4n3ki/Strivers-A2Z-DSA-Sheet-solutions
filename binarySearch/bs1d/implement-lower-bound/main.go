package main

import (
	"fmt"
)

func main() {
	arr1 := []int {1, 2, 2, 3}
	x1 := 2

	arr2 := []int {1, 2, 2, 4, 4, 4, 6, 7, 9}
	x2 := 4

	bruteforceApp(arr1, x1)
	bruteforceApp(arr2, x2)

	bs(arr1, x1)
	bs(arr2, x2)
}

func bs(arr []int, k int) {
	n := len(arr)
	index := binarySearch(arr, 0, n-1, k)
	fmt.Printf("arr: %v, index: %d\n", arr, index)
}

func binarySearch(arr []int, low, high, target int) int {
	index := high

	for low <= high {
		mid := (low+high)/2
		if arr[mid] >= target {
			high = mid - 1
			index = mid
		}else {
			low = mid + 1
		}
	}
	return index
}

func bruteforceApp(arr []int, k int) {
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] >= k {
			break
		}
	}
	fmt.Printf("arr: %v, index: %d\n", arr, i)
}