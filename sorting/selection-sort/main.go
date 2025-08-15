package main

import "fmt"

func main() {
	arr1 := []int {13,46,24,52,20,9}
	arr2 := []int {5,4,3,2,1}

	selectionSort(arr1)
	selectionSort(arr2)

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i+1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}