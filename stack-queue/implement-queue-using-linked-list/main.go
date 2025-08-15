package main

import (
	"fmt"
)

func main() {
	q := new(Queue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	q.Print()

	q.Dequeue()
	q.Print()

	fmt.Println(q.Top())
}

type Queue struct {
	front *Node
	end *Node
}

func (q *Queue) Enqueue(Data int) {
	if q.front == nil {
		q.front = &Node{Data, nil}
		q.end = q.front
	}else {
		q.end.next = &Node{Data, nil}
		q.end = q.end.next
	}
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	val := q.front.val
	next := q.front.next
	q.front.next = nil
	q.front = next
	return val, nil
}

func (q *Queue) Top() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	val := q.front.val
	return val, nil
}

func  (q *Queue) Size() int {
	cnt := 0
	curr := q.front
	for curr != nil {
		cnt++
		curr = curr.next
	}
	return cnt
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}
func (q *Queue) Print() {
	curr := q.front
	for curr != nil {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
	fmt.Println()
}

type Node struct {
	val int
	next *Node
}