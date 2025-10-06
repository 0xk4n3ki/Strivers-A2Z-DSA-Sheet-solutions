package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {1, 2, 3, 6, 2, 3, 4, 7, 8}
	n := 3
	fmt.Printf("ans: %v\n", partition(arr, n))

	arr = []int {1, 2, 3, 4, 5}
	n = 4
	fmt.Printf("ans: %v\n", partition(arr, n))
}

func partition(hand []int, n int) bool {
	if len(hand)%n != 0 {
		return false
	}

	freq := make(map[int]int)
	for _, card := range hand {
		freq[card]++
	}

	keys := make([]int, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, start := range keys {
		count := freq[start]
		if count == 0 {
			continue
		}

		for i := 0; i < n; i++ {
			card := start + i
			if freq[card] < count {
				return false
			}
			freq[card] -= count
		}
	}
	return true
}