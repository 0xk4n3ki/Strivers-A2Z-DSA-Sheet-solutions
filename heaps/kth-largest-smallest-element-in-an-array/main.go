package main

import "fmt"

func main() {
	arr := []int {1, 2, 6, 4, 5, 3}
	k := 3
	fmt.Println(smallest(arr, k), largest(arr, k))

	arr = []int {1, 2, 6, 4, 5}
	k = 3
	fmt.Println(smallest(arr, k), largest(arr, k))
}


// n length priority queue
func smallest(arr []int, k int) int {
	h := &MinHeap{}
	h.BuildHeap(arr)
	for i := 0; i < k-1; i++ {
		h.Pop()
	}
	return h.arr[0]
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

func largest(arr []int, k int) int {
	h := &MaxHeap{}
	h.BuildHeap(arr)
	for i := 0; i < k-1; i++ {
		h.Pop()
	}
	return h.arr[0]
}

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.arr = append([]int {}, arr...)
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	lastIndex := len(h.arr)-1
	for {
		l, r := 2*i+1, 2*i+2
		if l > lastIndex {
			return
		}
		indexToCompare := l
		if r <= lastIndex && h.arr[r] > h.arr[indexToCompare] {
			indexToCompare = r
		}
		if h.arr[i] < h.arr[indexToCompare] {
			h.arr[i], h.arr[indexToCompare] = h.arr[indexToCompare], h.arr[i]
			i = indexToCompare
		}else {
			return
		}
	}
}

func (h *MaxHeap) Pop() {
	if len(h.arr) == 0 {
		return
	}
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
}





// sorting method
func smallest1(arr []int, k int) int {
	quickSort(arr, 0, len(arr)-1)
	return arr[k-1]
}

func quickSort(arr []int, start, end int) {
	if start < end {
		pIndex := partition(arr, start, end)
		quickSort(arr, start, pIndex-1)
		quickSort(arr, pIndex+1, end)
	}
}

func partition(arr []int, start, end int) int {
	i, j := start, end
	pivot := start

	for i < j {
		for i < end && arr[i] <= arr[pivot] {
			i++
		}
		for j > start && arr[j] > arr[pivot] {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[pivot] = arr[pivot], arr[j]
	return j
}