package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}

	bruteforceApp(arr)
	optimalApp(arr)
}

func optimalApp(arr []int) {
	count := mergeSort(arr, 0, len(arr)-1)

	fmt.Printf("sorted arr: %v, count: %d\n", arr, count)
}

func bruteforceApp(arr []int) {
	n := len(arr)
	count := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] {
				count++
			}
		}
	}

	fmt.Printf("arr: %v, inversion count: %d\n", arr, count)
}

func mergeSort(arr []int, low, high int) int {
	count := 0
	if low >= high {
		return count
	}
	mid := (low + high) / 2
	count += mergeSort(arr, low, mid)
	count += mergeSort(arr, mid+1, high)
	count += merge(arr, low, mid, high)
	return count
}

func merge(arr []int, low, mid, high int) int {
	tmp := []int{}

	left, right := low, mid+1

	count := 0

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			tmp = append(tmp, arr[left])
			left++
		} else {
			count += (mid - left + 1)
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
	return count
}
