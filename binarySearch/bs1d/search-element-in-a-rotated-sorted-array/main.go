package main

import "fmt"

func main() {
	arr1 := []int{4, 5, 6, 7, 0, 1, 2, 3}
	k1 := 0

	arr2 := []int{4, 5, 6, 7, 0, 1, 2}
	k2 := 3

	bs(arr1, k1)
	bs(arr2, k2)
}

func bs(arr []int, target int) {
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			index = mid
			break
		}

		if arr[low] <= arr[mid] {
			if arr[low] <= target && target <= arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if arr[mid] <= target && target <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, target, index)
}
