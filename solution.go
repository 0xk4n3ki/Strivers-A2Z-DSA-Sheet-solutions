package main

import (
	"fmt"
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

func main() {
	boards := []int {5, 5, 5, 5}
	k := 2
	fmt.Printf("boards: %v, ans: %d\n", boards, calc(boards, k))

	boards = []int {10, 20, 30, 40}
	k = 2
	fmt.Printf("boards: %v, ans: %d\n", boards, calc(boards, k))
}

func calc(boards []int, k int) int {
	low, high := calcTimelimits(boards)

	for low <= high {
		mid := (low+high)/2

		currK := calcTime(boards, mid)
		fmt.Printf("curr: %d, low: %d, mid: %d, high: %d\n", currK, low, mid, high)

		if currK <= k {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return low
}

func calcTime(boards []int, time int) int {
	if len(boards) == 0 {
		return 0
	}
	paintersNum := 0
	sum := 0

	for _, i := range boards {
		if sum + i > time {
			sum = i
			paintersNum += 1
		} else {
			sum += i
		}
	}
	if sum > 0 {
		paintersNum++
	}
	return paintersNum
}

func calcTimelimits(boards []int) (int, int) {
	min := 0
	max := 0
	for _, i := range boards {
		max += i
		if min < i {
			min = i
		}
	}
	return min, max
}