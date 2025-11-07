package main

import (
	"fmt"
	"sort"
)

// [OLD] https://leetcode.com/problems/largest-odd-number-in-string/description/
// [OLD] https://takeuforward.org/data-structure/quick-sort-algorithm/
// [OLD] https://takeuforward.org/data-structure/segregate-even-and-odd-nodes-in-linkedlist
// [NEW] https://takeuforward.org/arrays/painters-partition-problem/
// [NEW] https://leetcode.com/problems/roman-to-integer/description/
// [NEW] https://takeuforward.org/data-structure/combination-sum-ii-find-all-unique-combinations/
// [NEW] https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/
// [NEW] https://leetcode.com/problems/count-primes/description/
// [NEW] https://takeuforward.org/plus/dsa/problems/postfix-to-prefix-conversion
// [NEW] https://leetcode.com/problems/sum-of-subarray-minimums/description/
// [NEW] https://leetcode.com/problems/minimum-window-substring/


func main () {
	arr := []int {10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Printf("arr: %v, target: %d, ans: %v\n", arr, target, calc(arr, target))

	arr = []int {2, 5, 2, 1, 2}
	target = 5
	fmt.Printf("arr: %v, target: %d, ans: %v\n", arr, target, calc(arr, target))
}

func calc(arr []int, target int) [][]int {
	sort.Ints(arr)
	ans := [][]int {}
	curr := []int {}
	backtracking(0, target, arr, &ans, &curr)
	return ans
}

func backtracking(index, target int, arr []int, ans *[][]int, curr *[]int) {
	if target == 0 {
		*ans = append(*ans, append([]int {}, *curr...));
		return
	}

	for i := index; i < len(arr); i++ {
		if i > index && arr[i] == arr[i-1] {
			continue
		}
		if arr[i] > target {
			break
		}
		*curr = append(*curr, arr[i])
		backtracking(i+1, target-arr[i], arr, ans, curr)
		*curr = (*curr)[:len(*curr)-1]
	}
}