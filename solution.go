package main

import (
	"fmt"
)

// [OLD] https://takeuforward.org/data-structure/median-of-row-wise-sorted-matrix/
// [OLD] https://takeuforward.org/data-structure/peak-element-in-array/
// [OLD] https://takeuforward.org/data-structure/count-reverse-pairs/
// [NEW] https://takeuforward.org/arrays/kth-missing-positive-number/
// [NEW] https://leetcode.com/problems/string-to-integer-atoi/description/
// [NEW] https://takeuforward.org/data-structure/power-set-print-all-the-possible-subsequences-of-the-string/
// [NEW] https://takeuforward.org/data-structure/word-search-leetcode/
// [NEW] https://leetcode.com/problems/subsets/description/
// [NEW] https://takeuforward.org/data-structure/check-for-balanced-parentheses/
// [NEW] https://leetcode.com/problems/remove-k-digits/
// [NEW] https://leetcode.com/problems/longest-repeating-character-replacement/description/

func main() {
	arr := []int {1,3,2,3,1}
	ans := count(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {3,2,1,4}
	ans = count(arr)
	fmt.Printf("ans: %d\n", ans)
}

func count(arr []int) int {
	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, low, high int) int {
	cnt := 0
	if low < high {
		mid := (low+high)/2
		cnt += mergeSort(arr, low, mid)
		cnt += mergeSort(arr, mid+1, high)
		cnt += countPairs(arr, low, mid, high)
		merge(arr, low, mid, high)
	}
	return cnt
}

func countPairs(arr []int, low, mid, high int) int {
	right := mid + 1
	cmt := 0
	for i := low; i <= mid; i++ {
		for right <= high && arr[i] > 2*arr[right] {
			right++
		}
		cmt += (right-(mid+1))
	}
	return cmt
}

func merge(arr []int, low, mid, high int) {
	left, right := low, mid + 1
	tmp := []int {}

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			tmp = append(tmp, arr[left])
			left++
		}else {
			tmp = append(tmp, arr[right])
			right++
		}
	}

	for left <= mid + 1 {
		tmp = append(tmp, arr[left])
		left++
	}
	for right <= high {
		tmp = append(tmp, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = tmp[i-low]
	}
}