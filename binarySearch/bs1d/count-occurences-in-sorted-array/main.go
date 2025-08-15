package main

import (
	"fmt"
)

func main() {
	arr1 := []int {2, 2, 3, 3, 3, 3, 4}
	x1 := 3

	arr2 := []int {1, 1, 2, 2, 2, 2, 2, 3}
	x2 := 2

	bs(arr1, x1)
	bs(arr2, x2)
}

func bs(arr []int, target int) {
	first := firstOccurence(arr, target)
	if first == -1 {
		fmt.Printf("arr: %v, target: %d, count: %d\n", arr, target, -1)
		return
	}
	last := lastOccurence(arr, target)
	count := last - first + 1
	fmt.Printf("arr: %v, target: %d, count: %d\n", arr, target, count)
}

func firstOccurence(arr []int, target int) int {
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] == target {
			index = mid
			high = mid - 1
		}else if arr[mid] > target {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return index
}

func lastOccurence(arr []int, target int) int {
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] == target {
			index = mid
			low = mid + 1
		}else if arr[mid] > target {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return index
}