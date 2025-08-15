package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {1, 2, 2}
	ans := subSum(arr)
	fmt.Printf("ans: %v\n", ans)
}

func subSum(arr []int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sort.Ints(arr)
	subRec(arr, 0, &sub, &ans)

	return ans
}

func subRec(arr []int, i int, sub *[]int, ans *[][]int) {
	tmp := []int {}
	tmp = append(tmp, *sub...)
	*ans = append(*ans, tmp)

	for j := i; j < len(arr); j++ {
		if j > i && arr[j] == arr[j-1] {
			continue
		}
		*sub = append(*sub, arr[j])
		subRec(arr, j+1, sub, ans)
		*sub = (*sub)[:len(*sub)-1]


	}
}