package main

import "fmt"

func main() {
	arr := []int {2,0,2,1,1,0}
	fmt.Printf("before: %v ", arr)
	basicSumApp(arr)
	fmt.Printf("after: %v\n", arr)


	arr = []int {2,0,1,2,0,1,0,1,1,0}
	fmt.Printf("before: %v ", arr)
	threePointerApp(arr)
	fmt.Printf("after: %v\n", arr)
}

func basicSumApp(arr []int) {
	zeros, ones, twos := 0, 0, 0
	for _, i := range arr {
		if i == 0 {
			zeros++
		}else if i == 1 {
			ones++
		}else {
			twos++
		}
	}

	index := 0
	for i := 0; i < zeros; i++ {
		arr[index] = 0
		index++
	}
	for i := 0; i < ones; i++ {
		arr[index] = 1
		index++
	}
	for i := 0; i < twos; i++ {
		arr[index] = 2
		index++
	}
}

func threePointerApp(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		}else if arr[mid] == 1 {
			mid++
		}else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}