package main

import (
	"fmt"
)

func main() {
	arr1 := []int{4, 1, 7, 9, 3}
	arr2 := []int{4, 6, 2, 5, 7, 9, 1, 3}

	quickSort(arr1, 0, len(arr1)-1)
	quickSort(arr2, 0, len(arr2)-1)

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pIndex := partition(arr, low, high)

		quickSort(arr, low, pIndex-1)
		quickSort(arr, pIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low, high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}

		for arr[j] > pivot && j >= low+1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
