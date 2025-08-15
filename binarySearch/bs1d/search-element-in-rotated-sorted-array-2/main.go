package main

import "fmt"

func main() {
	arr := []int {7, 8, 1, 2, 3, 3, 3, 4, 5, 6}
	k1 := 3
	k2 := 10

	bs(arr, k1)
	bs(arr, k2)

	arr = []int {3, 1, 2, 3, 3, 3, 3}
	k := 1
	bs(arr, k)
}

func bs(arr []int, k int) {
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low + high)/2

		if arr[mid] == k {
			index = k
			break
		}

		// without this, the 3rd case won't work
		if arr[low] == arr[mid] && arr[mid] == arr[high] {
			low = low + 1
			high = high - 1
		}

		if arr[low] <= arr[mid] {
			if arr[low] <= k && k <= arr[mid] {
				high = mid - 1
			}else {
				low = mid + 1
			}
		}else {
			if arr[mid] <= k && k <= arr[high] {
				low = mid + 1
			}else {
				high = mid - 1
			}
		}
	}
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, k, index)
}