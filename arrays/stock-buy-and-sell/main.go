package main

import "fmt"

func main() {
	arr := []int {7,1,5,3,6,4}
	profit := basicApp(arr)
	fmt.Printf("arr: %v, profit: %d\n", arr, profit)


	arr = []int {3,6,3,7,2,8,23,8,1}
	profit = optimalApp(arr)
	fmt.Printf("arr: %v, profit: %d\n", arr, profit)
}

func basicApp(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			diff := arr[j] - arr[i]

			if diff > max {
				max = diff
			}
		}
	}
	return max
}

func optimalApp(arr []int) int {
	maxProfit, min := 0, arr[0]
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if maxProfit < arr[i] - min {
			maxProfit = arr[i] - min
		}
	}
	return maxProfit
}