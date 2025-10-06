package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2, arr3 := []int{1, 2, 3, 4}, []int{-4, -3}, []int{-5, -3, 1, 2, 3, 4}
	arr := []*Node{arr2ll(arr1), arr2ll(arr2), arr2ll(arr3)}
	ans := brute(arr)
	printll(ans)
	ans2 := merge(arr)
	printll(ans2)
}


func merge(heads []*Node) *Node {
	h := &HeapNode{}

	for _, head := range heads {
		if head != nil {
			h.Insert(head)
		}
	}
	
	dummy := &Node{}
	tail := dummy

	for len(h.arr) > 0 {
		smallest := h.ExtractMin()
		tail.next = smallest
		tail = tail.next

		if smallest.next != nil {
			h.Insert(smallest.next)
		}
	}
	return dummy.next
}



type HeapNode struct {
	arr []*Node
}

func parent(i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *HeapNode) Insert(node *Node) {
	h.arr = append(h.arr, node)
	i := len(h.arr) - 1

	for i > 0 && h.arr[parent(i)].val > h.arr[i].val {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

func (h *HeapNode) Heapify(i int) {
	smallest := i
	l, r := left(i), right(i)
	n := len(h.arr)

	if l < n && h.arr[l].val < h.arr[smallest].val {
		smallest = l
	}
	if r < n && h.arr[r].val < h.arr[smallest].val {
		smallest = r
	}

	if smallest != i {
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		h.Heapify(smallest)
	}
}

func (h *HeapNode) ExtractMin() *Node {
	if len(h.arr) == 0 {
		return nil
	}
	if len(h.arr) == 1 {
		min := h.arr[0]
		h.arr = h.arr[:0]
		return min
	}
	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.Heapify(0)

	return min
}

func (h *HeapNode) GetMin() *Node {
	if len(h.arr) == 0 {
		return nil
	}
	return h.arr[0]
}




func brute(heads []*Node) *Node {
	ansArr := []int{}

	for _, head := range heads {
		tmp := head
		for tmp != nil {
			ansArr = append(ansArr, tmp.val)
			tmp = tmp.next
		}
	}

	sort.Ints(ansArr)
	return arr2ll(ansArr)
}

type Node struct {
	val  int
	next *Node
}

func arr2ll(arr []int) *Node {
	head := &Node{-1, nil}
	tmp := head
	for _, i := range arr {
		tmp.next = &Node{i, nil}
		tmp = tmp.next
	}
	return head.next
}

func printll(head *Node) {
	tmp := head
	for tmp != nil {
		fmt.Printf("%d ", tmp.val)
		tmp = tmp.next
	}
	fmt.Println()
}
