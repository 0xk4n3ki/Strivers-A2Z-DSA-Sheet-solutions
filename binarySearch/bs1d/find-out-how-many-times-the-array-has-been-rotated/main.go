package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {4, 5, 6, 7, 0, 1, 2, 3}
	bs(arr)

	arr = []int {10, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	bs(arr)
}

func bs(arr []int) {
	low, high, index, minNum := 0, len(arr)-1, -1, math.MaxInt

	for low <= high {
		mid := (low + high)/2
		
		if arr[low] <= arr[high] {
			if minNum > arr[low] {
				minNum = arr[low]
				index = low
			}
			break
		}

		if arr[low] <= arr[mid] {
			if minNum > arr[low] {
				minNum = arr[low]
				index = low

			}
			low = mid + 1
		}else {
			if arr[mid] < minNum {
				minNum = arr[mid]
				index = mid
			}
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, minNum: %d, times of rotation: %d\n", arr, minNum, index)
}