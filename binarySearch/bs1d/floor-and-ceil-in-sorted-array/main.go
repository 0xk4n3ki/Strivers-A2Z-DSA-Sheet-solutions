package main

import "fmt"

func main() {
	arr1 := []int {3, 4, 4, 7, 8, 10}
	x1 := 5
	x2 := 8

	floorFunc(arr1, x1)
	floorFunc(arr1, x2)

	ceilFunc(arr1, x1)
	ceilFunc(arr1, x2)
}

func ceilFunc(arr []int, target int){
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] >= target {
			high = mid - 1
			index = mid
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("arr: %v, target: %d, ceil: %d\n", arr, target, arr[index])
}

func floorFunc(arr []int, target int) {
	low, high, index := 0, len(arr)-1, -1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] <= target {
			low = mid + 1
			index = mid
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, target: %d, floor: %d\n", arr, target, arr[index])
}