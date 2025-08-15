package main

import (
	"fmt"
)

func main() {
	arr := []int {0, 3, 4, 7, 10, 9}
	k := 4

	method(arr, k)
	bs(arr, k)

	arr = []int {4, 2, 1, 3, 6}
	k = 2

	method(arr, k)
	bs(arr, k)
}

func bs(arr []int, k int) {
	n := len(arr)
	quickSort(arr, 0, n-1)

	low, high := 1, arr[n-1] - arr[0]
	for low <= high {
		mid := (low + high)/2

		if canWePlace(arr, k, mid) {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, cows: %d, min dist: %d\n", arr, k, high)
}

func method(arr []int, k int) {
	quickSort(arr, 0, len(arr)-1)
	n := len(arr)
	limit := arr[n-1] - arr[0]
	var i int
	for i = 1; i <= limit; i++ {
		if !canWePlace(arr, k, i) {
			break
		}
	}
	fmt.Printf("arr: %v, cows: %d, min dist: %d\n", arr, k, i-1)
}

func canWePlace(arr []int, k, dist int) bool {
	cntCows, last := 1, arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] - last >= dist {
			cntCows++
			last = arr[i]
		}
		if cntCows >= k {
			return true
		}
	}
	return false
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pIndex := partition(arr, low, high)
		quickSort(arr, low, pIndex)
		quickSort(arr, pIndex + 1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]

	i, j := low, high

	for i < j {
		for arr[i] <= pivot && i <= high - 1 {
			i++
		}
		for arr[j] > pivot && j >= low + 1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[low] = arr[low], arr[j]
	return j
}