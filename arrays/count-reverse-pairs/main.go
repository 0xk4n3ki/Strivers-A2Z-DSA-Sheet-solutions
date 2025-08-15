package main

import "fmt"

func main() {
	arr := []int {1, 3, 2, 3, 1}

	bruteforceApp(arr)
	usingMergeSort(arr)
}

func usingMergeSort(arr []int) {
	count := mergeSort(arr, 0, len(arr)-1)
	fmt.Printf("arr: %v, count: %d\n", arr, count)
}


func bruteforceApp(arr []int) {
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > 2 * arr[j] {
				count++
			}
		}
	}

	fmt.Printf("arr: %v, count: %d\n", arr, count)
}


func mergeSort(arr []int, low, high int) int {
	count := 0
	if low >= high {
		return count
	}
	mid := (low+high)/2
	count += mergeSort(arr, low, mid)
	count += mergeSort(arr, mid+1, high)
	count += countPairs(arr, low, mid, high)
	merge(arr, low, mid, high)
	return count
}

func countPairs(arr []int, low, mid, high int) int {
	right := mid + 1
	count := 0

	for i := low; i <= mid; i++ {
		for right <= high && arr[i] > 2 * arr[right] {
			right++
		}
		count += right - (mid + 1)
	}
	return count
}

func merge(arr []int, low, mid, high int) {
	tmp := []int {}

	left, right := low, mid+1

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			tmp = append(tmp, arr[left])
			left++
		}else {
			tmp = append(tmp, arr[right])
			right++
		}
	}

	for left <= mid {
		tmp = append(tmp, arr[left])
		left++
	}
	for right <= high {
		tmp = append(tmp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = tmp[i-low]
	}
}