package main

import "fmt"

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

	len := llLen(head)

	fmt.Printf("len: %d\n", len)

	printll(head)
}

func llLen(head *Node) int {
	len := 0
	for head != nil {
		len++
		head = head.next
	}

	return len
}


func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Println(head.val)
}