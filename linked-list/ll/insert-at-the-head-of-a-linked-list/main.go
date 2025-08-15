package main

import "fmt"

type Node struct {
	val int
	next *Node
}



func main() {
	head := &Node{12, nil}
	cur := head
	
	cur.next = &Node{5, nil}
	cur = cur.next

	cur.next = &Node{8, nil}
	cur = cur.next

	cur.next = &Node{7, nil}
	cur = cur.next

	printLL(head)

	head = insertFront(head, 100)

	printLL(head)
}

func insertFront(head *Node, n int) *Node {
	new := &Node{n, head}
	return new
}

func printLL(head *Node) {
	cur := head

	for cur.next != nil {
		fmt.Printf("%d->", cur.val)
		cur = cur.next
	}
	fmt.Printf("%d\n", cur.val)
}


func insertEnd(head *Node, n int) {
	cur := head

	new := &Node{n, nil}

	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new
}