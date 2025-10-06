package main

import (
	"fmt"
)

func main() {
	arr := []rune {'A', 'A', 'B', 'A', 'B', 'B'}
	n := 2
	fmt.Println(leastInterval(arr, n))

	arr = []rune {'A', 'C', 'A', 'B', 'D', 'B'}
	n = 1
	fmt.Println(leastInterval(arr, n))
}

func leastInterval(tasks []rune, n int) int {
	freq := make(map[rune]int)
	for _, task := range tasks {
		freq[task]++
	}

	h := &MaxHeap{}
	for _, v := range freq {
		h.Insert(v)
	}

	time := 0
	for !h.IsEmpty() {
		cycle := n+1
		temp := []int {}
		i := 0

		for i < cycle && !h.IsEmpty() {
			cnt, _ := h.ExtractMax()
			cnt--
			if cnt > 0 {
				temp = append(temp, cnt)
			}
			time++
			i++
		}

		h.PushAll(temp)
		if !h.IsEmpty() {
			time += (cycle - i)
		}
	}
	return time
}




type MaxHeap struct {
	arr []int
}

func parent(i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

func (h *MaxHeap) Insert(key int) {
	h.arr = append(h.arr, key)
	i := len(h.arr) - 1

	for i > 0 && h.arr[parent(i)] < h.arr[i] {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

func (h *MaxHeap) Heapify(i int) {
	l, r := left(i), right(i)
	largest := i
	size := len(h.arr)

	if l < size && h.arr[l] > h.arr[largest] {
		largest = l
	}
	if r < size && h.arr[r] > h.arr[largest] {
		largest = r
	}

	if largest != i {
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		h.Heapify(largest)
	}
}

func (h *MaxHeap) ExtractMax() (int, bool) {
	size := len(h.arr)
	if size == 0 {
		return 0, false
	}

	max := h.arr[0]
	h.arr[0] = h.arr[size-1]
	h.arr = h.arr[:size-1]
	if len(h.arr) > 0 {
		h.Heapify(0)
	}
	return max, true
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.arr) == 0
}

func (h *MaxHeap) PushAll(values []int) {
	for _, v := range values {
		h.Insert(v)
	}
}