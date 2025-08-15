package main

import "fmt"

func main() {
	arr := []int {4, 7, 9, 10}
	k := 1

	bruteforce(arr, k)
	bs(arr, k)

	arr = []int {4, 7, 9, 10}
	k = 4
	
	bruteforce(arr, k)
	bs(arr, k)
}

func bs(arr []int, k int) {
	n := len(arr)
	low, high := 0, n - 1

	for low <= high {
		mid := (low + high)/2

		missingNums := arr[mid] - (mid + 1)
		
		if missingNums < k {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	ans := k + high + 1

	fmt.Printf("arr: %v, k: %d, missing number: %d\n", arr, k, ans)
}

func bruteforce(arr []int, k int) {
	dup := k
	for i := 0; i < len(arr); i++ {
		if arr[i] <= k {
			k++
		}else {
			break
		}
	}
	fmt.Printf("arr: %v, k: %d, missing number: %d\n", arr, dup, k)
}