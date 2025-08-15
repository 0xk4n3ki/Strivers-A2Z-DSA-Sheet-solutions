package main

import (
	"fmt"
)

func main() {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Print()
	s.Pop()
	s.Print()
	fmt.Println(s.Top())
}

type Node struct {
	val int
	next *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(data int) {
	tmp := &Node{data, s.head}
	s.head = tmp
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.head.val
	s.head = s.head.next
	return val, nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.head.val, nil
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Size() int {
	count := 0
	for curr := s.head; curr != nil; curr = curr.next {
		count++
	}
	return count
}

func (s *Stack) Print() {
	for curr := s.head; curr != nil; curr = curr.next {
		fmt.Print(curr.val, " ")
	}
	fmt.Println()
}
 























// type Stack struct {
// 	head *Node
// 	end *Node
// }

// func (s *Stack) Print() {
// 	curr := s.head
// 	for curr != nil {
// 		fmt.Print(curr.val, " ")
// 		curr = curr.next
// 	}
// 	fmt.Println()
// }

// func (s *Stack) Push(data int) {
// 	tmp := &Node{data, nil}
// 	if s.IsEmpty() {
// 		s.head = tmp
// 		s.end = tmp
// 	}else {
// 		s.end.next = tmp
// 		s.end = tmp
// 	}
// }

// func (s *Stack) Pop() (int, error) {
// 	if s.IsEmpty() {
// 		return 0, fmt.Errorf("stack is empty")
// 	}
// 	curr := s.head
// 	for curr.next != s.end {
// 		curr = curr.next
// 	}

// 	val := s.end.val
// 	curr.next = nil
// 	s.end = curr

// 	return val, nil
// }

// func (s *Stack) Top() (int, error) {
// 	if s.IsEmpty() {
// 		return 0, fmt.Errorf("stack is empty")
// 	}
// 	return s.end.val, nil
// }

// func (s *Stack) Size() int {
// 	n := 0
// 	curr := s.head
// 	for curr != nil {
// 		curr = curr.next
// 		n++
// 	}
// 	return n
// }

// func (s *Stack) IsEmpty() bool {
// 	return s.head == nil
// }

