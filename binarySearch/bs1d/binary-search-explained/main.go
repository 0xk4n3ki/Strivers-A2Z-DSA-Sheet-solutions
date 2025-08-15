package main

import "fmt"

func main() {
	// arr := []int {3, 4, 6, 7, 9, 12, 16, 17}
	arr := []int {1, 3, 5, 7, 9}
	target := 8
	
	binarySearch(arr, target)
	index := recApp(arr, 0, len(arr)-1, target)
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, target, index)
}

func recApp(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}

	mid := (low + high)/2

	if arr[mid] == target {
		return mid
	}else if arr[mid] > target {
		return recApp(arr, low, mid - 1, target)
	}
	return recApp(arr, mid + 1, high, target)
}

func binarySearch(arr []int, target int) {
	low, high := 0, len(arr)-1

	index := -1

	for low <= high {
		mid := (high + low)/2
		if arr[mid] == target {
			index = mid
			break
		}else if arr[mid] < target {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, target, index)
}