package main

import (
	"fmt"
	"math"
)

type MinHeap struct {
	array []int
}

func parent(index int) int {
	return (index-1)/2
}

func left(index int) int {
	return 2*index+1
}

func right(index int) int {
	return 2*index+2
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)

	index := len(h.array)-1

	for index > 0 && h.array[parent(index)] > h.array[index] {
		h.array[parent(index)], h.array[index] = h.array[index], h.array[parent(index)]
		index = parent(index)
	}
}

func (h *MinHeap) Heapify(index int) {
	r, l := right(index), left(index)

	smallest := index
	lastIndex := len(h.array)-1

	if l <= lastIndex && h.array[l] < h.array[smallest] {
		smallest = l
	}
	if r <= lastIndex && h.array[r] < h.array[smallest] {
		smallest = r
	}

	if smallest != index {
		h.array[smallest], h.array[index] = h.array[index], h.array[smallest]
		h.Heapify(smallest)
	}
}

func (h *MinHeap) GetMin() int {
	return h.array[0]
}

func (h *MinHeap) ExtractMin() int {
	size := len(h.array)
	if size == 0 {
		return math.MinInt
	}

	if size == 1 {
		min := h.array[0]
		h.array = h.array[:size-1]
		return min
	}

	min := h.array[0]
	h.array[0] = h.array[size-1]

	h.Heapify(0)
	return min
}

func (h *MinHeap) DecreaseKey(i, val int) {
	h.array[i] = val

	for i != 0 && h.array[parent(i)] > h.array[i] {
		h.array[parent(i)], h.array[i] = h.array[i], h.array[parent(i)]
		i = parent(i)
	}
}

func (h *MinHeap) Delete(i int) {
	h.DecreaseKey(i, math.MinInt)
	h.ExtractMin()
}

func main() {
	m := MinHeap{}
	buildHeap := []int {10, 20, 30, 5, 40, 1, 45, 95, 75, 80}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	fmt.Println("min value:", m.GetMin())

	m.Insert(-1)
	fmt.Println("min value:", m.GetMin())
	fmt.Println(m)

	m.DecreaseKey(3, -2)
	fmt.Println("min value:", m.GetMin())
	fmt.Println(m)

	m.ExtractMin()
	fmt.Println("min value:", m.GetMin())
	fmt.Println(m)

	m.Delete(0)
	fmt.Println("min value:", m.GetMin())

	fmt.Println(m)
}