package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {10, 1, 2, 7, 6, 1, 5}
	target := 8

	ans := comb(arr, target)
	fmt.Printf("arr: %v, target: %d, ans: %v\n", arr, target, ans)


	ans2 := comb2(arr, target)
	fmt.Printf("arr: %v, target: %d, ans: %v\n", arr, target, ans2)
}

// sol 1
func comb(arr []int, target int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sort.Ints(arr)

	backtracking(0, arr, target, &sub, &ans)
	return ans
}

func backtracking(index int, arr []int, target int, sub *[]int, ans *[][]int) {
	if target == 0 {
		tmp := []int {}
		tmp = append(tmp, *sub...)
		*ans = append(*ans, tmp)
		return
	}

	if index >= len(arr) || target < 0 {
		return
	}

	*sub = append(*sub, arr[index])
	backtracking(index+1, arr, target-arr[index], sub, ans)

	*sub = (*sub)[:len(*sub)-1]

	for index < len(arr)-1 && arr[index] == arr[index+1] {
		index++
	}
	backtracking(index+1, arr, target, sub, ans)
}


// sol2
func comb2(arr []int, target int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sort.Ints(arr)

	backtracking2(0, arr, target, &sub, &ans)
	return ans
}

func backtracking2(index int, arr []int, target int, sub *[]int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int {}, *sub...))
		return
	}

	for i := index; i < len(arr); i++ {
		if i > index && arr[i] == arr[i-1] {
			continue
		}

		if arr[i] > target {
			break
		}

		*sub = append(*sub, arr[i])
		backtracking(i+1, arr, target-arr[i], sub, ans)
		*sub = (*sub)[:len(*sub)-1]
	}
}