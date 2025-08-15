package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 1, 9}
	n, k := 5, 10
	len := basicApp(arr, n, k)
	fmt.Printf("n: %d, k: %d, arr: %v, len: %d\n", n, k, arr, len)

	arr = []int{2, 5, 4, 4, 2, 2, 2, 2, 6, 2}
	n, k = 10, 8
	len = usingTwoArrays(arr, n, k)
	fmt.Printf("n: %d, k: %d, arr: %v, len: %d\n", n, k, arr, len)

	arr = []int{2, 5, 4, 4, 2, 2, 2, 2, 6, 2}
	n, k = 10, 8
	len = hashingApp(arr, n, k)
	fmt.Printf("n: %d, k: %d, arr: %v, len: %d\n", n, k, arr, len)

	// arr = []int {2, 5, 4, 4, 2, 2, 2, 2, 6, 2}
	arr = []int{2, 3, 5, 1, 9}
	n, k = 5, 10
	len = twoPointerApp(arr, n, k)
	fmt.Printf("n: %d, k: %d, arr: %v, len: %d\n", n, k, arr, len)
}

func basicApp(arr []int, n, k int) int {
	len := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for l := i; l <= j; l++ {
				sum += arr[l]
			}
			if sum == k {
				len = max(len, j-i+1)
			}
		}
	}
	return len
}

func usingTwoArrays(arr []int, n, k int) int {
	len := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == k {
				len = max(len, j-i+1)
			}
		}
	}
	return len
}

// best in case of positive and negative
func hashingApp(arr []int, n, k int) int {
	m := map[int]int{}
	sum, len := 0, 0
	for i := 0; i < n; i++ {
		sum += arr[i]

		if sum == k {
			len = max(len, i+1)
		}

		rem := sum - k

		// If rem is in the map, update maxLen
		if idx, ok := m[rem]; ok {
			l := i - idx
			len = max(len, l)
		}

		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
	}
	return len
}

// best only for positive and zeroes
func twoPointerApp(arr []int, n, k int) int {
	len := 0
	sum := 0
	i := 0
	for j := 0; j < n; j++ {
		sum += arr[j]

		for i <= j && sum > k {
			sum -= arr[i]
			i++
		}

		if sum == k {
			len = max(len, j-i+1)
		}
	}
	return len
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
