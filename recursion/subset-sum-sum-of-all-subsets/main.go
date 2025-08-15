package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {5, 2, 1}
	ans := subsetSum(arr)
	fmt.Printf("arr: %v, ans: %v\n", arr, ans)
}

func subsetSum(arr []int) []int {
	ans:= []int {}

	backtracking(0, arr, 0, &ans)
	
	sort.Ints(ans)
	return ans
}

func backtracking(index int, arr []int, sum int, ans *[]int) {
	if index == len(arr) {
		*ans = append(*ans, sum)
		return
	}

	backtracking(index+1, arr, sum+arr[index], ans)
	backtracking(index+1, arr, sum, ans)
}