package main

import "fmt"

func main() {
	arr1 := []int {13,46,24,52,20,9}
	arr2 := []int {5,4,3,2,1}

	bubbleSort(arr1)
	bubbleSort(arr2)

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func bubbleSort(arr []int) {
	for i := len(arr)-1; i>=0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}