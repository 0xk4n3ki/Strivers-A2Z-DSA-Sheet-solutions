package main

import "fmt"

func main() {
	arr1 := []int {13,46,24,52,20,9}
	arr2 := []int {5,4,3,2,1}

	insertionSort(arr1)
	insertionSort(arr2)

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}