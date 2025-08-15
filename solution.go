package main

import (
	"fmt"
)

// [OLD] https://takeuforward.org/arrays/count-subarray-sum-equals-k/
// [OLD] https://takeuforward.org/data-structure/quick-sort-algorithm/
// [OLD] https://leetcode.com/problems/isomorphic-strings/description/
// [NEW] https://takeuforward.org/binary-search/koko-eating-bananas/
// [NEW] https://leetcode.com/problems/sort-characters-by-frequency/description/
// [NEW] https://leetcode.com/problems/generate-parentheses/description/
// [NEW] https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/
// [NEW] https://leetcode.com/problems/single-number/description/
// [NEW] https://takeuforward.org/data-structure/check-for-balanced-parentheses/
// [NEW] https://takeuforward.org/data-structure/trapping-rainwater/
// [NEW] https://leetcode.com/problems/count-number-of-nice-subarrays/description/

func main() {
	arr := []int {1, 1, 2, 1, 1}
	k := 3
	ans := count(arr, k)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {2, 4, 6}
	k = 1
	ans = count(arr, k)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {2,2,2,1,2,2,1,2,2,2}
	k = 2
	ans = count(arr, k)
	fmt.Printf("ans: %d\n", ans)
}

func count(arr []int, k int) int {
	return nice(arr, k)-nice(arr, k-1)
}

func nice(arr []int, k int) int {
	if k < 0 {
		return 0
	}
	cnt := 0
	i, j := 0, 0
	ans := 0

	for j < len(arr) {
		cnt += arr[j] % 2

		for cnt > k {
			cnt -= arr[i] % 2
			i++
		}
		if cnt <= k {
			ans += j-i+1
		}
		j++
	}
	return ans
}