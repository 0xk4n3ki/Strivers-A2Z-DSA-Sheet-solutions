package main

import "fmt"

func main() {
	n, k := 7, 2
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	side := "right"

	fmt.Printf("before: %v ", arr)
	rightRotateReverseAlgo(arr, n, k)
	fmt.Printf("after %s rotation by %d places: %v\n", side, k, arr)

	n, k = 6, 2
	arr = []int{3, 7, 8, 9, 10, 11}
	side = "left"
	fmt.Printf("before: %v ", arr)
	leftRotateReverseAlgo(arr, n, k)
	fmt.Printf("after %s rotation by %d places: %v\n", side, k, arr)
}

// approach 1 : using a tmp array

// func leftRotate(arr []int, n, k int) {
// 	if n == 0 {
// 		return
// 	}
// 	k = k % n

// 	tmp := []int{}
// 	tmp = append(tmp, arr[:k]...)

// 	for i := 0; i < n-k; i++ {
// 		arr[i] = arr[i+k]
// 	}
// 	for i := 0; i < k; i++ {
// 		arr[n-k+i] = tmp[i]
// 	}
// }

// func rightRotate(arr []int, n, k int) {
// 	if n == 0 {
// 		return
// 	}
// 	k = k % n

// 	tmp := []int {}
// 	tmp = append(tmp, arr[:n-k]...)
	
// 	for i := 0; i < k; i++ {
// 		arr[i] = arr[n-k+i]
// 	}
// 	for i := 0; i < n-k; i++ {
// 		arr[i+k] = tmp[i]
// 	}
// }

// approach 2 : using reversal algorithm

func rightRotateReverseAlgo(arr []int, n, k int) {
	reverse(arr, 0, n-k-1)
	reverse(arr, n-k, n-1)
	reverse(arr, 0, n-1)
}

func leftRotateReverseAlgo(arr []int, n, k int) {
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
	reverse(arr, 0, n-1)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}