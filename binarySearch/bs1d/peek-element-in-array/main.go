package main

import "fmt"

func main() {
	arr1 := []int {1,2,3,4,5,6,7,8,5,1}
	arr2 := []int {1, 2, 1, 3, 5, 6, 4}
	arr3 := []int {1, 2, 3, 4, 5}
	arr4 := []int {5, 4, 3, 2, 1}

	bruteforceApp(arr1)
	bruteforceApp(arr2)
	bruteforceApp(arr3)
	bruteforceApp(arr4)

	bs(arr1)
	bs(arr2)
	bs(arr3)
	bs(arr4)
}

func bs(arr []int) {
	index := bsFunc(arr)
	fmt.Printf("arr: %v, index: %d, peak element: %d\n", arr, index, arr[index])
}

func bsFunc(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}
	if arr[0] > arr[1] {
		return 0
	}
	if arr[n-1] > arr[n-2] {
		return n-1
	}

	low, high := 1, n-2

	for low <= high {
		mid := (low + high)/2

		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid - 1] {
			return mid
		}else if arr[mid] < arr[mid + 1] {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return -1
}

func bruteforceApp(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if (i == 0 || arr[i-1] < arr[i]) && (i == n-1 || arr[i] > arr[i+1]) {
			fmt.Printf("arr: %v, index: %d, peak element: %d\n", arr, i, arr[i])
			break
		}
	}
}