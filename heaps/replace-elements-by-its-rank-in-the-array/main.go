package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int {20, 15, 26, 2, 98, 6}
	ans := replace(arr)
	fmt.Println(ans)

	arr = []int {1, 5, 8, 15, 8, 25, 9}
	ans = replace(arr)
	fmt.Println(ans)
}

// O(nlogn)
func replace(arr []int) []int {
	h := &MinHeap{}
	h.BuildHeap(arr)

	hashMap := map[int]int {}
	tmp := 1

	for len(h.arr) > 0 {
		val := h.arr[0]
		h.Pop()
		if _, ok := hashMap[val]; !ok {
			hashMap[val] = tmp
			tmp++
		}
	}


	ans := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ans[i] = hashMap[arr[i]]
	}
	return ans
}

type MinHeap struct {
	arr []int
}

func (h *MinHeap) BuildHeap(arr []int) {
	h.arr = append([]int {}, arr...)
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MinHeap) heapifyDown(i int) {
	lastIndex := len(h.arr)-1
	for {
		l, r := 2*i+1, 2*i+2
		if l > lastIndex {
			return
		}
		indexToCompare := l
		if r <= lastIndex && h.arr[r] < h.arr[indexToCompare] {
			indexToCompare = r
		}

		if h.arr[i] > h.arr[indexToCompare] {
			h.arr[i], h.arr[indexToCompare] = h.arr[indexToCompare], h.arr[i]
			i = indexToCompare
		}else {
			return
		}
	}
}

func (h *MinHeap) Pop() {
	if len(h.arr) == 0 {
		return 
	}

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
}


// O(n) + O(nlogn) + O(n)
func replace2(arr []int) []int {
	hashMap := map[int]int {}
	tmp := 1
	brr := append([]int {}, arr...)

	sort.Ints(brr)
	for i := 0; i < len(arr); i++ {
		if hashMap[brr[i]] == 0 {
			hashMap[brr[i]] = tmp
			tmp++
		}
	}

	ans := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ans[i] = hashMap[arr[i]]
	}
	return ans
}


// O(n^2)
func replace1(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		hashMap := map[int]struct{} {}
		for j := 0; j < len(arr); j++ {
			if arr[j] < arr[i] {
				hashMap[arr[j]] = struct{}{}
			}
		}
		ans[i] = len(hashMap)+1
	}
	return ans
}