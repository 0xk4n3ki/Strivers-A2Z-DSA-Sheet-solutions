package main

import "fmt"

func main() {
	arr1 := []int {1, 2, 4, 7}
	x1 := 6
	x2 := 2

	bruteforceApp(arr1, x1)
	bruteforceApp(arr1, x2)

	bs(arr1, x1)
	bs(arr1, x2)
}

func bs(arr []int, target int) {
	index := -1
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low+high)/2

		if arr[mid] >= target {
			index = mid
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, target, index)
}


func bruteforceApp(arr []int, target int) {
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] >= target {
			break
		}
	}
	fmt.Printf("arr: %v, target: %d, index: %d\n", arr, target, i)
}