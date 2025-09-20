package main

import (
	"fmt"
	"math"
)

// Heap Interface
type Heap interface {
	Insert(key int)
	Extract() int
	Peek() int
	BuildHeap(arr []int)
	Delete(i int)
	Size() int
}

// MinHeap
type MinHeap struct {
	array []int
}

func (h *MinHeap) parent(i int) int {
	return (i-1)/2
}
func (h *MinHeap) left(i int) int {
	return 2*i + 1
}
func (h *MinHeap) right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) Size() int {
	return len(h.array)
}

func (h *MinHeap) Peek() int {
	if len(h.array) == 0 {
		return math.MinInt
	}
	return h.array[0]
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array)-1)
}

func (h *MinHeap) Extract() int {
	if h.Size() == 0 {
		return math.MinInt
	}

	root := h.array[0]
	last := h.array[h.Size()-1]
	h.array[0] = last
	h.array = h.array[:h.Size()-1]
	h.heapifyDown(0)
	return root
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.array[h.parent(i)] > h.array[i] {
		h.array[h.parent(i)], h.array[i] = h.array[i], h.array[h.parent(i)]
		i = h.parent(i)
	}
}

func (h *MinHeap) heapifyDown(i int) {
	lastIndex := h.Size() - 1

	for {
		l, r := h.left(i), h.right(i)
		if l > lastIndex {
			break
		}

		childToCompare := l
		if r <= lastIndex && h.array[r] < h.array[l] {
			childToCompare = r
		}

		if h.array[i] > h.array[childToCompare] {
			h.array[i], h.array[childToCompare] = h.array[childToCompare], h.array[i]
			i = childToCompare
		}else {
			break
		}
	}
}

func (h *MinHeap) BuildHeap(arr []int) {
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MinHeap) DecreaseKey(i, val int) {
	h.array[i] = val
	h.heapifyUp(i)
}

func (h *MinHeap) Delete(i int) {
	h.DecreaseKey(i, math.MinInt)
	h.Extract()
}




// MaxHeap
type MaxHeap struct {
	array []int
}

func (h *MaxHeap) parent(i int) int {
	return (i-1)/2
}
func (h *MaxHeap) left(i int) int {
	return 2*i+1
}
func (h *MaxHeap) right(i int) int {
	return 2*i+2
}

func (h *MaxHeap) Size() int {
	return len(h.array)
}

func (h *MaxHeap) Peek() int {
	if h.Size() == 0 {
		return math.MaxInt
	}
	return h.array[0]
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(h.Size()-1)
}

func (h *MaxHeap) Extract() int {
	if h.Size() == 0 {
		return math.MaxInt
	}

	root := h.array[0]
	last := h.array[h.Size()-1]
	h.array[0] = last
	h.array = h.array[:h.Size()-1]
	h.heapifyDown(0)
	return root
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 && h.array[h.parent(i)] < h.array[i] {
		h.array[i], h.array[h.parent(i)] = h.array[h.parent(i)], h.array[i]
		i = h.parent(i)
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	lastIndex := h.Size()-1

	for {
		l, r := h.left(i), h.right(i)
		if l > lastIndex {
			break
		}

		childToCompare := l
		if r <= lastIndex && h.array[r] > h.array[l] {
			childToCompare = r
		}

		if h.array[i] < h.array[childToCompare] {
			h.array[i], h.array[childToCompare] = h.array[childToCompare], h.array[i]
			i = childToCompare
		}else {
			break
		}
	}
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) IncreaseKey(i, val int) {
	h.array[i] = val
	h.heapifyUp(i)
}

func (h *MaxHeap) Delete(i int) {
	h.IncreaseKey(i, math.MaxInt)
	h.Extract()
}


func main() {
	fmt.Println("MinHeap:")
	m := &MinHeap{}
	m.BuildHeap([]int {10, 20, 30, 5, 40, 1, 45, 95, 75, 80})
	fmt.Println("Heap:", m.array)
	fmt.Println("Peek:", m.Peek())

	m.Insert(-2)
	fmt.Println("After Insert(-2):", m.array)
	fmt.Println("Extract:", m.Extract())
	fmt.Println("After Extract:", m.array)



	fmt.Println("MaxHeap:")
	M := &MaxHeap{}
	M.BuildHeap([]int {10, 20, 30, 5, 40, 1, 45, 95, 75, 80})
	fmt.Println("Heap:", M.array)
	fmt.Println("Peek:", M.Peek())

	M.Insert(200)
	fmt.Println("After Insert(200):", M.array)
	fmt.Println("Extract:", M.Extract())
	fmt.Println("After Extract:", M.array)
}