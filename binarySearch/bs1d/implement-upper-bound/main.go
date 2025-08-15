package main

import "fmt"

func main() {
	arr1 := []int {1, 2, 2, 3}
	x1 := 2

	arr2 := []int {3, 5, 8, 9, 15, 19}
	x2 := 6

	bruteforceApp(arr1, x1)
	bruteforceApp(arr2, x2)

	optimalApp(arr1, x1)
	optimalApp(arr2, x2)
}

func optimalApp(arr []int, target int) {
	index := len(arr)
	low, high := 0, index-1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] <= target {
			low = mid + 1
		}else {
			high = mid - 1
			index = mid
		}
	}
	fmt.Printf("arr: %v, index: %d\n", arr, index)
}

func bruteforceApp(arr []int, target int) {
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] > target {
			break
		}
	}
	fmt.Printf("arr: %v, index: %d\n", arr, i)
}