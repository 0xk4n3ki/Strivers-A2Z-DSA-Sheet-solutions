package main

import "fmt"

type Node struct {
	val int
	next *Node
	back *Node
}

func main() {
	head := &Node{2, nil, nil}
	head.next = &Node{4, nil, head}
	head.next.next = &Node{5, nil, head.next}
	head.next.next.next = &Node{7, nil, head.next.next}

	fmt.Printf("%p, %v\n", head, head)
	fmt.Printf("%p, %v\n", head.next, head.next)
	fmt.Printf("%p, %v\n", head.next.next, head.next.next)
	fmt.Printf("%p, %v\n", head.next.next.next, head.next.next.next)
}