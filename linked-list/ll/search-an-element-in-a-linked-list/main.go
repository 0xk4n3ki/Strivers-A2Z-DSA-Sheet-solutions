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

	searchll(head, 2)
	searchll(head, 373)
}

func searchll(head *Node, n int) {
	cur := head

	for cur != nil {
		if cur.val == n {
			fmt.Printf("n: %d, ans: %t\n", n, true)
			return
		}
		cur = cur.next
	}
	fmt.Printf("n: %d, ans: %t\n", n, false)
}

func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Println(head.val)
}