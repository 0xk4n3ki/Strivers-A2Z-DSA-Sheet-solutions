package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {1, 1, 1, 2, 2}
	target := 4
	ans := combFunc(arr, target)
	fmt.Printf("ans: %v\n", ans)
}

func combFunc(arr []int, target int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sort.Ints(arr)
	combRec(arr, 0, target, &sub, &ans)
	return ans
}

func combRec(arr []int, i, target int, sub *[]int, ans *[][]int) {
	if target == 0 {
		tmp := []int {}
		tmp = append(tmp, *sub...)
		*ans = append(*ans, tmp)
		return
	}

	for j := i; j < len(arr); j++ {
		if j > i && arr[j] == arr[j-1] {
			continue
		}
		if arr[i] > target {
			break
		}

		*sub = append(*sub, arr[i])
		combRec(arr, j+1, target-arr[j], sub, ans)
		*sub = (*sub)[:len(*sub)-1]
	}
}