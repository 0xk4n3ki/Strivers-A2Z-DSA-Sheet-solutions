package main

import "fmt"

func main() {
	arr1 := []int {13,46,24,52,20,9}
	arr2 := []int {5,4,3,2,1}

	recInsertion(arr1, 0, len(arr1))
	recInsertion(arr2, 0, len(arr2))

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func recInsertion(arr []int, i, len int) {
	if i == len  {
		return
	}
	for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
		arr[j], arr[j-1] = arr[j-1], arr[j]
	}
	recInsertion(arr, i+1, len)
}