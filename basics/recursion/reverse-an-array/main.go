package main

import "fmt"

func main() {
	arr1 := []int {5, 4, 3, 2, 1}
	fmt.Printf("reverse of %v: ", arr1)
	fmt.Printf("%v\n", revArr(arr1, 0, len(arr1)-1))

	arr2 := []int {10, 20, 30, 40, 50}
	fmt.Printf("reverse of %v: ", arr2)
	fmt.Printf("reverse %v\n", revArr(arr2, 0, len(arr2)-1))
}

func revArr(arr []int, i, l int) []int {
	if i >= l {
		return arr
	}
	arr[i], arr[l] = arr[l], arr[i]
	return revArr(arr, i+1, l-1)
}