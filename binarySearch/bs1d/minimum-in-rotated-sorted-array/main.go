package main

import (
	"fmt"
	"math"
)

func main() {
	arr1 := []int {4,5,6,7,0,1,2,3}
	bs(arr1)

	arr2 := []int {3, 4, 5, 1, 2}
	bs(arr2)
}

func bs(arr []int) {
	low, high, minNum := 0, len(arr)-1, math.MaxInt

	for low <= high {
		mid := (low + high)/2

		if arr[low] <= arr[high] {
			minNum = min(arr[low], minNum)
			break
		}

		if arr[low] <= arr[mid] {
			minNum = min(arr[low], minNum)
			low = mid + 1
		}else {
			minNum = min(arr[mid], minNum)
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, min: %d\n", arr, minNum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}