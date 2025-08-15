package main

import (
	"fmt"
)

func main() {
	arr := []int {-2,1,-3,4,-1,2,1,-5,4}
	sum := basicApp(arr)
	fmt.Printf("arr: %v, sum: %d\n", arr, sum)


	arr = []int {-2,-3,4,-1,-2,1,5,-3}
	sum, subArr := optimalApp(arr)
	fmt.Printf("arr: %v, sum: %d, subArr: %v\n", arr, sum, subArr)
}

func basicApp(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]

			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func optimalApp(arr []int) (int, []int) {
	max := -10000
	sum := 0
	ansStart, ansEnd, start := -1, -1, 0
	for i := 0; i < len(arr); i++ {
		if sum == 0 {
			start = i
		}
		sum += arr[i]

		if sum > max {
			max = sum
			ansStart, ansEnd = start, i
		}

		if sum < 0 {
			sum = 0
		}
	}
	return max, arr[ansStart:ansEnd+1]
}