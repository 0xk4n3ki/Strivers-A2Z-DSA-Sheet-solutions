package main

import (
	"fmt"
)

func main() {
	arr, k := []int {1, 2, 1, 2, 3}, 2
	num, ans := brute(arr, k)
	fmt.Printf("ans: %d, %v\n", num, ans)
	num = optimal(arr, k)
	fmt.Printf("ans: %d\n", num)

	arr, k = []int {1, 2, 1, 3, 4}, 3
	num, ans = brute(arr, k)
	fmt.Printf("ans: %d, %v\n", num, ans)
	num = optimal(arr, k)
	fmt.Printf("ans: %d\n", num)
}

func optimal(arr []int, k int) int {
	return atmost(arr, k) - atmost(arr, k-1)
}

func atmost(arr []int, k int) int {
	num := 0
	left, right:= 0, 0
	numMap := map[int]int {}

	for right < len(arr) {
		numMap[arr[right]]++

		for len(numMap) > k {
			numMap[arr[left]]--
			if numMap[arr[left]] == 0 {
				delete(numMap, arr[left])
			}
			left++
		}
		num += right - left + 1
		right++
	}
	return num
}

func brute(arr []int, k int) (int, [][]int) {
	n := len(arr)
	cnt, ans := 0, [][]int {}

	for i := 0; i < n; i++ {
		numMap := map[int]int {}
		for j := i; j < n; j++ {
			numMap[arr[j]]++

			if len(numMap) == k {
				cnt++
				ans = append(ans, arr[i:j+1])
			}

			if len(numMap) > k {
				break
			}
		}
	}
	return cnt, ans
}