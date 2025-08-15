package main

import "fmt"

func main() {
	arr := []int {1,2,4,5}
	n := 5
	m := sumApp(arr, n)
	fmt.Printf("missing number in %v is: %d\n", arr, m)

	arr = []int {1,2,3,4,5,7,8}
	n = 8
	m = xorApp(arr, n)
	fmt.Printf("missing number in %v is: %d\n", arr, m)

	arr = []int {1,2,3,4,5,6,7,8,10,11,12}
	n = 12
	m = hashingApp(arr, n)
	fmt.Printf("missing number in %v is: %d\n", arr, m)
}

func sumApp(arr []int, n int) int {
	sum := int(n*(n+1)/2)

	newSum := 0
	for _, i := range arr {
		newSum += i
	}
	return sum - newSum
}

func xorApp(arr []int, n int) int {
	xor1, xor2 := 0, 0

	for i := 0; i < n-1; i++ {
		xor1 ^= i+1
		xor2 ^= arr[i]
	}
	xor1 ^= n
	
	return xor1 ^ xor2
}

func hashingApp(arr []int, n int) int {
	arr1 := make([]int, n+1)

	for _, i := range arr {
		arr1[i]++
	}

	for i := 1; i <= n; i++ {
		if arr1[i] == 0 {
			return i
		}
	}
	return -1
}