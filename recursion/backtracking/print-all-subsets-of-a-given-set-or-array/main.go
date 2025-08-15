package main

import (
	"fmt"
	"sort"
)

// https://www.geeksforgeeks.org/dsa/backtracking-to-find-all-subsets/

func main() {

	arr := []int {1, 2, 3}

	res := backtracking(arr)

	fmt.Printf("res: %v\n", res)
}

func backtracking(arr []int) [][]int {
	subset := []int {}
	ans := [][]int {}

	sort.Ints(arr)

	subsetRec(0, arr, &ans, &subset)

	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})

	return ans
}

func subsetRec(i int, arr []int, ans *[][]int, subset *[]int) {
	if i == len(arr) {
		temp := make([]int, len(*subset))
		copy(temp, *subset)
		*ans = append(*ans, temp)
		return
	}

	*subset = append(*subset, arr[i])
	subsetRec(i+1, arr, ans, subset)

	*subset = (*subset)[:len(*subset)-1]
	subsetRec(i+1, arr, ans, subset)
}