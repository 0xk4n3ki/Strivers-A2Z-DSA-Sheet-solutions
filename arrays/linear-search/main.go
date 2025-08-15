package main

import "fmt"

func main() {
	arr1 := []int {1, 2, 3, 4, 5}
	n1 := 3
	ind1 := linear(arr1, n1)
	fmt.Printf("index of %d in %v is: %d\n", n1, arr1, ind1)

	// arr2 := []int {5, 4, 3, 2, 1}
	// n2 := 5
	// ind2 := optimal(arr2, n2)
	// fmt.Printf("index of %d in %v is: %d\n", n2, arr2, ind2)
}

func linear(arr []int, n int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}