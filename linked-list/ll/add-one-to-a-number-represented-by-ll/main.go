package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {1, 2, 3}
	head := arr2ll(arr)
	printll(head)

	head = addOne(head)

	printll(head)


	arr = []int {9, 9}
	head = arr2ll(arr)
	printll(head)
	head = addOne(head)
	printll(head)


	arr = []int {6, 9}
	head = arr2ll(arr)
	printll(head)
	head = addOne(head)
	printll(head)
}

func addOne(head *Node) *Node {
	rev := revll(head)

	curr := rev
	carry := 1

	for curr != nil {
		sum := curr.val + carry
		curr.val = sum % 10
		carry = sum / 10

		if carry == 0 {
			break
		}

		if curr.next == nil && carry != 0 {
			curr.next = &Node{carry, nil}
			carry = 0
			break
		}

		curr = curr.next
	}

	return revll(rev)
}

func revll(head *Node) *Node {
	curr := head
	var prev *Node = nil
	
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
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