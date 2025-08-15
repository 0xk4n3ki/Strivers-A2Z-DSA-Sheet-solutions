package main

import "fmt"

func main() {
	arr1 := []int {13,46,24,52,20,9}
	arr2 := []int {5,4,3,2,1}

	recBubble(arr1, len(arr1))
	recBubble(arr2, len(arr2))

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func recBubble(arr []int, len int) {
	if len == 1 {
		return
	}
	didSwap := false
	for j := 0; j < len - 1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			didSwap = true
		}
	}
	if !didSwap {
		return
	}
	recBubble(arr, len-1)
}