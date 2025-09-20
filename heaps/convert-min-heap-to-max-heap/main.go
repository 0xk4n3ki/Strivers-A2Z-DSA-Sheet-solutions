package main

import "fmt"

func main() {
	arr := []int {10, 20, 30, 21, 23}
	convert(arr)
	fmt.Println(arr)

	arr = []int {-5, -4, -3, -2, -1}
	convert(arr)
	fmt.Println(arr)
}

func convert(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(arr, i)
	}
}

func heapifyDown(arr []int, i int) {
	lastIndex := len(arr)-1

	for {
		l, r := 2*i+1, 2*i+2
		if l > lastIndex {
			return
		}

		indexToCompare := l
		if r <= lastIndex && arr[r] > arr[indexToCompare] {
			indexToCompare = r
		}

		if arr[i] < arr[indexToCompare] {
			arr[i], arr[indexToCompare] = arr[indexToCompare], arr[i]
			i = indexToCompare
		}else {
			break
		}
	}
}