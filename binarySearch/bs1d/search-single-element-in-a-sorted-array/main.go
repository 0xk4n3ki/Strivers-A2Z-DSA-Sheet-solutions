package main

import "fmt"

func main() {
	arr := []int {1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6}
	arr2 := []int {1, 1, 3, 5, 5}

	bs(arr)
	bs(arr2)
}

func bs(arr []int) {
	n := len(arr)
	if n == 1 {
		fmt.Printf("arr: %v, ans: %d\n", arr, arr[0])
		return
	}
	if arr[0] != arr[1] {
		fmt.Printf("arr: %v, ans: %d\n", arr, arr[0])
		return
	}
	if arr[n-1] != arr[n-2] {
		fmt.Printf("arr: %v, ans: %d\n", arr, arr[n-1])
		return
	}

	low, high, ans := 1, len(arr)-2, -1

	for low <= high {
		mid := (low + high)/2
		if arr[mid] != arr[mid-1] && arr[mid] != arr[mid+1] {
			ans = arr[mid]
			break
		}

		if (mid % 2 != 0 && arr[mid] == arr[mid-1]) || (mid % 2 == 0 && arr[mid] == arr[mid+1]) {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, ans: %d\n", arr, ans)
}