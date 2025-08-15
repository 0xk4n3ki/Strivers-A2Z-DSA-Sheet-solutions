package main

import (
	"container/heap"
	"fmt"
)


type Section struct {
	length float64
	index int
}

type MaxHeap []Section

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	if h[i].length != h[j].length {
		return h[i].length > h[j].length
	}
	return h[i].index > h[j].index
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Section))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func main() {
	arr := []int {1, 2, 3, 4, 5}
	k := 4
	bruteforce(arr, k)
	betterApp(arr, k)
	optimalApp(arr, k)

	arr = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k = 1
	bruteforce(arr, k)
	betterApp(arr, k)
	optimalApp(arr, k)
}

func optimalApp(arr []int, k int) {
	n := len(arr)

	low, high := float64(0), float64(0)

	for i := 0; i < n-1; i++ {
		if high < float64(arr[i+1] - arr[i]) {
			high = float64(arr[i+1] - arr[i])
		}
	}

	diff := 0.000001
	for high - low > diff {
		mid := (low + high)/2.0

		cnt := calcStations(arr, mid)
		if cnt > k {
			low = mid
		}else {
			high = mid
		}
	}
	fmt.Printf("arr: %d, k: %d, dist: %f\n", arr, k, high)
}

func calcStations(arr []int, dist float64) int {
	n, cnt := len(arr), 0

	for i := 1; i < n; i++{
		numInBetween := int(float64((arr[i] - arr[i-1]))/dist)

		if float64(arr[i] - arr[i-1]) == dist * float64(numInBetween) {
			numInBetween--
		}
		cnt += numInBetween
	}
	return cnt
}

func betterApp(arr []int , k int) {
	n := len(arr)
	howMany := make([]int, n-1)

	pq := &MaxHeap{}

	for i := 0; i < n-1; i++ {
		diff := float64(arr[i+1] - arr[i])
		heap.Push(pq, Section{length: diff, index: i})
	}

	for i := 0; i < k; i++ {
		sec := heap.Pop(pq).(Section) // pop the maximum distand and its index
		idx := sec.index
		howMany[idx]++

		initialDist := float64(arr[idx+1] - arr[idx])
		newLen := initialDist/float64(howMany[idx]+1)
		heap.Push(pq, Section{length: newLen, index: idx})
	}

	fmt.Printf("arr: %d, k: %d, ans: %f\n", arr, k, (*pq)[0].length)
}

func bruteforce(arr []int, k int) {
	n := len(arr)
	howMany := make([]float32, n-1)

	for station := 1; station <= k; station++ {
		maxSection, maxIndex := float32(-1), -1

		for i := 0; i < n-1; i++ {
			diff := float32(arr[i+1] - arr[i])

			sectionLen := diff / float32(howMany[i] + 1)

			if sectionLen > maxSection {
				maxSection = sectionLen
				maxIndex = i
			}
		}
		howMany[maxIndex]++
	}

	maxAns := float32(-1)
	for i := 0; i < n-1; i++ {
		diff := float32(arr[i+1] - arr[i])
		sectionLen := diff/float32(howMany[i] + 1)

		if maxAns < sectionLen {
			maxAns = sectionLen
		}
	}

	fmt.Printf("arr: %v, k: %d, ans: %f\n", arr, k, maxAns)
}