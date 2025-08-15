package main

import "fmt"

func main() {
	arr1 := []int{4, 2, 1, 6, 7}
	arr2 := []int{3, 2, 8, 5, 1, 4, 23}

	mergeSort(arr1, 0, len(arr1)-1)
	mergeSort(arr2, 0, len(arr2)-1)

	fmt.Printf("sorted arr1: %v\n", arr1)
	fmt.Printf("sorted arr2: %v\n", arr2)
}

func mergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)

	merge(arr, low, mid, high)
}

func merge(arr []int, low, mid, high int) {
	tmp := []int{}
	left, right := low, mid+1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			tmp = append(tmp, arr[left])
			left++
		} else {
			tmp = append(tmp, arr[right])
			right++
		}
	}
	// if elements on left are still left
	for left <= mid {
		tmp = append(tmp, arr[left])
		left++
	}
	// if elements on right are still left
	for right <= high {
		tmp = append(tmp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = tmp[i-low]
	}
}
