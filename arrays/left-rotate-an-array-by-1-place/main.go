package main

import "fmt"

func main() {
	arr1 := []int {1,2,3,4,5}
	arr2 := []int {3}

	fmt.Printf("before: %v ", arr1)
	leftRotate(arr1)
	fmt.Printf("after left rotate by one place: %v\n", arr1)

	fmt.Printf("before: %v ", arr2)
	leftRotate(arr2)
	fmt.Printf("after left rotate by one place: %v\n", arr2)
}

func leftRotate(arr []int) {
	tmp := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = tmp
}