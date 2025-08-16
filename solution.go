package main

import (
	"fmt"
)

// [OLD] https://takeuforward.org/data-structure/count-occurrences-in-sorted-array/
// [OLD] https://takeuforward.org/data-structure/reverse-words-in-a-string/
// [OLD] https://takeuforward.org/data-structure/find-intersection-of-two-linked-lists/
// [NEW] https://takeuforward.org/arrays/painters-partition-problem/
// [NEW] https://leetcode.com/problems/roman-to-integer/description/
// [NEW] https://takeuforward.org/plus/dsa/problems/check-if-there-exists-a-subsequence-with-sum-k
// [NEW] https://takeuforward.org/data-structure/word-search-leetcode/
// [NEW] https://leetcode.com/problems/single-number/description/
// [NEW] https://takeuforward.org/data-structure/implement-min-stack-o2n-and-on-space-complexity/
// [NEW] https://leetcode.com/problems/maximal-rectangle/description/
// [NEW] https://leetcode.com/problems/binary-subarrays-with-sum/description/

func main() {
	nums := []int {1, 0, 1, 0, 1}
	goal := 2
	ans := binary(nums, goal)
	fmt.Printf("ans: %d\n", ans)

	nums = []int {0, 0, 0, 0, 0}
	goal = 0
	ans = binary(nums, goal)
	fmt.Printf("ans: %d\n", ans)
}

func binary(arr []int, goal int) int {
	return sum(arr, goal)-sum(arr, goal-1)
}


func sum(arr []int, goal int) int {
	if goal < 0 {
		return 0
	}
	cnt := 0
	sum := 0
	i, j := 0, 0

	for j < len(arr) {
		sum += arr[j]%2
		for sum > goal {
			sum -= arr[i]
			i++
		}
		cnt += j-i+1
		j++
	}
	return cnt
}