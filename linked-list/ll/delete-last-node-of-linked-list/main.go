package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

func main() {
	head := &Node{372, nil}
	head.next = &Node{1, nil}
	head.next.next = &Node{2, nil}
	head.next.next.next = &Node{3, nil}
	head.next.next.next.next = &Node{4, nil}
	head.next.next.next.next.next = &Node{5, nil}

	printll(head)

	head = deleteLastNode(head)

	printll(head)
}

func deleteLastNode(head *Node) *Node {
	cur := head

	if cur == nil || cur.next == nil {
		return nil
	}

	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil

	return head
}

func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Println(head.val)
}