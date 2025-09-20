package main

import "fmt"

func main() {
	arr := []int {6, 5, 3, 2, 8, 10, 9}
	k := 3
	sort(arr, k)
	fmt.Println(arr)

	arr = []int {1, 4, 5, 2, 3, 6, 7, 8, 9, 10}
	k = 2
	sort(arr, k)
	fmt.Println(arr)
}

func sort(arr []int, k int) {
	h := &MinHeap{}
	for i := 0; i < k; i++ {
		h.Insert(arr[i])
	}

	var i int
	for i = k; i < len(arr); i++ {
		h.Insert(arr[i])
		arr[i-k] = h.arr[0]
		h.Pop()
	}

	for len(h.arr) != 0 {
		arr[i-k] = h.arr[0]
		h.Pop()
		i++
	}
}

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.heapifyUp(len(h.arr)-1)
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.arr[i] < h.arr[(i-1)/2] {
		h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
		i = (i-1)/2
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

func (h *MinHeap) heapifyDown(i int) {
	lastIndex := len(h.arr)-1
	for {
		l, r := 2*i+1, 2*i+2
		if l > lastIndex {
			return
		}
		indexToCompare := l
		for r <= lastIndex && h.arr[r] < h.arr[indexToCompare] {
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