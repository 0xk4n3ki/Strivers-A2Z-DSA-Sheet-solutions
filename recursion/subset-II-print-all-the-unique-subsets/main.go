package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {4, 4, 4, 1, 4}
	ans := subset(arr)

	fmt.Printf("arr: %v, ans: %v\n", arr, ans)

	arr = []int {4, 4, 4, 1, 4}
	ans = subset2(arr)

	fmt.Printf("arr: %v, ans: %v\n", arr, ans)
}


// sol1
func subset2(arr []int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sort.Ints(arr)

	backtracking2(0, arr, &sub, &ans)

	return ans
}

func backtracking2(index int, arr []int, sub *[]int, ans *[][]int) {
	*ans = append(*ans, append([]int {}, *sub...))

	for i := index; i < len(arr); i++ {
		if i != index && arr[i] == arr[i-1] {
			continue
		}

		*sub = append(*sub, arr[i])
		backtracking2(i+1, arr, sub, ans)

		*sub = (*sub)[:len(*sub)-1]
	}
}




// sol2
func subset(arr []int) [][]int {
	ans := [][]int {}
	sub := []int {}

	sort.Ints(arr)
	
	backtracking(0, arr, &sub, &ans)

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

func backtracking(index int, arr []int, sub *[]int, ans *[][]int) {
	if index >= len(arr) {
		*ans = append(*ans, append([]int {}, *sub...))
		return
	}

	*sub = append(*sub, arr[index])
	backtracking(index+1, arr, sub, ans)
	*sub = (*sub)[:len(*sub)-1]


	for index < len(arr)-1 && arr[index] == arr[index+1] {
		index++
	}

	backtracking(index+1, arr, sub, ans)
}