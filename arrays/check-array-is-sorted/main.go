package main

import "fmt"

func main() {
	arr1 := []int {1, 2, 5, 4, 6, 7, 8}
	arr2 := []int {5,6,7,8,9}

	fmt.Printf("%v is sorted: %v\n", arr1, checkSort(arr1))
	fmt.Printf("%v is sorted: %v\n", arr2, checkSort(arr2))
}

// func checkSort(arr []int) bool {
// 	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
// 		if arr[i] > arr[i+1] || arr[j] < arr[j-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

func checkSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}