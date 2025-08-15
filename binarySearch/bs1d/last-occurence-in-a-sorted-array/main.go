package main

import "fmt"

func main() {
	arr := []int {3, 4, 13, 13, 13, 20, 40}
	x1 := 13
	x2 := 60

	bs(arr, x1)
	bs(arr, x2)
}

func bs(arr []int, target int) {
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] <= target {
			low = mid + 1
			if arr[mid] == target {
				index = mid
			}
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, target, index)
}