package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {1, 3, 2, 4}
	head := arr2ll(arr)
	printll(head)

	head = reversell(head)
	printll(head)

	head = inplaceIterative(head)
	printll(head)

	head = inplaceRecursion(head)
	printll(head)
}


func inplaceRecursion(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	newHead := inplaceRecursion(head.next)

	front := head.next

	head.next = nil
	front.next = head

	return newHead
}



func inplaceIterative(head *Node) *Node {
	if head == nil {
		return nil
	}

	var prev *Node = nil

	for head != nil {
		tmp := head.next
		head.next = prev

		prev = head
		head = tmp
	}

	return prev
}


func reversell(head *Node) *Node {
	if head == nil {
		return nil
	}

	rev := &Node{head.val, nil}
	head = head.next

	for head != nil {
		rev = &Node{head.val, rev}
		head = head.next
	}

	return rev
}



func lenll(head *Node) int {
	n := 0

	for head != nil {
		n++
		head = head.next
	}
	return n
}

func arr2ll(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{arr[0], nil}
	cur := head

	for i := 1; i < len(arr); i++ {
		tmp := &Node{arr[i], nil}
		cur.next = tmp

		cur = tmp
	}

	return head
}

func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Printf("%d\n", head.val)
}