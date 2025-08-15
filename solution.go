package main

import (
	"fmt"
)

// [OLD] https://takeuforward.org/binary-search/find-peak-element-ii
// [OLD] https://takeuforward.org/arrays/rearrange-array-elements-by-sign/
// [OLD] https://takeuforward.org/linked-list/delete-the-middle-node-of-the-linked-list
// [NEW] https://takeuforward.org/arrays/capacity-to-ship-packages-within-d-days/
// [NEW] https://leetcode.com/problems/roman-to-integer/description/
// [NEW] https://leetcode.com/problems/count-good-numbers/description/
// [NEW] https://takeuforward.org/data-structure/palindrome-partitioning/
// [NEW] https://leetcode.com/problems/minimum-bit-flips-to-convert-number/description/
// [NEW] https://takeuforward.org/data-structure/implement-min-stack-o2n-and-on-space-complexity/
// [NEW] https://leetcode.com/problems/sum-of-subarray-ranges/description/
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
	i, j := 0, 0
	cnt := 0
	ans := 0

	for j < len(arr) {
		cnt += arr[j] % 2

		for cnt > k {
			cnt -= arr[i]%2
			i++
		}

		if cnt <= k {
			ans += j-i+1
		}
		j++
	}
	return ans
}