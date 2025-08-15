package main

import "fmt"

func main() {
	arr := []int {3, 1, 2, 4}
	k := 6

	bruteforceApp(arr, k)

	// arr = []int {1, 2, 3}
	arr = []int {1, -1, 1, 1}
	k = 1

	betterApp(arr, k)
}

func betterApp(arr []int, k int) {
	preSum, count := 0, 0
	
	sumMap := make(map[int]int)
	sumMap[0] = 1

	for i := 0; i < len(arr); i++ {
		preSum += arr[i]

		rem := preSum - k

		if _, ok := sumMap[rem]; ok {
			count += sumMap[rem]
		}

		sumMap[preSum] += 1
	}

	fmt.Printf("arr: %v\nbetter App, count: %d\n", arr, count)
}



// doesn't work when array contains negative numbers
func bruteforceApp(arr []int, k int) {
	count := 0

	sum := 0
	startingIndex := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		// fmt.Printf("startingindex: %d, sum: %d\n", startingIndex, sum)
		if sum == k {
			count++
			startingIndex++
			i = startingIndex
			sum = 0
		}else if sum > k {
			startingIndex++
			i = startingIndex
			sum = 0
		}
	}
	fmt.Printf("arr: %v\ncount: %d\n", arr, count)
}