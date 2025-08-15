package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {3, 1, 2}
	ans := subSum(arr)
	fmt.Printf("ans: %v\n", ans)
}

func subSum(arr []int) []int {
	ans := []int {}

	subRec(arr, 0, 0, &ans)
	sort.Ints(ans)
	return ans
}

func subRec(arr []int, i, sum int, ans *[]int) {
	if i == len(arr) {
		*ans = append(*ans, sum)
		return
	}

	subRec(arr, i+1, sum+arr[i], ans)
	subRec(arr, i+1, sum, ans)
}