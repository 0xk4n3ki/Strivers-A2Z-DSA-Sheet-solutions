package main

import (
	"fmt"
)

func main() {
	arr := []int {4, 0, -1, 3, 5, 3, 6, 8}
	ans := slidingWindows(arr, 3)
	fmt.Printf("ans: %v\n", ans)
}

func slidingWindows(arr []int, k int) []int {
	ans := []int {}
	qe := new(DeQueue)

	for i := 0; i < len(arr); i++ {
		if !qe.IsEmpty() && qe.topFront() <= i-k {
			qe.popFront()
		}

		for !qe.IsEmpty() && arr[qe.topBack()] < arr[i] {
			qe.popBack()
		}

		qe.pushBack(i)

		if i >= k-1 {
			ans = append(ans, arr[qe.topFront()])
		}
	}
	return ans
}




type DeQueue struct {
	items []int
}

func (dq *DeQueue) pushFront(data int) {
	dq.items = append([]int {data}, dq.items...)
}

func (dq *DeQueue) pushBack(data int) {
	dq.items = append(dq.items, data)
}

func (dq *DeQueue) IsEmpty() bool {
	return dq.Size() == 0
}

func (dq *DeQueue) Size() int {
	return len(dq.items)
}

func (dq *DeQueue) topFront() int {
	if dq.IsEmpty() {
		return -1
	}
	return dq.items[0]
}

func (dq *DeQueue) topBack() int {
	if dq.IsEmpty() {
		return -1
	}
	return dq.items[dq.Size()-1]
}

func (dq *DeQueue) popFront() {
	if dq.IsEmpty() {
		return
	}
	dq.items = dq.items[1:]
}

func (dq *DeQueue) popBack() {
	if dq.IsEmpty() {
		return
	}
	dq.items = dq.items[:dq.Size()-1]
}
