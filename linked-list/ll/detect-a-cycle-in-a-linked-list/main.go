package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	head := &Node{1, nil}
	head.next = &Node{2, nil}
	head.next.next = &Node{3, nil}
	head.next.next.next = &Node{4, nil}
	head.next.next.next.next = &Node{9, nil}
	head.next.next.next.next.next = &Node{9, nil}
	head.next.next.next.next.next.next = &Node{5, nil}

	printll(head)
	bruteforce(head)
	optimal(head)

	head.next.next.next.next.next.next = &Node{5, head.next.next}
	bruteforce(head)
	optimal(head)
}

func optimal(head *Node) {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			fmt.Printf("loop detected\n")
			return
		}
	}
	fmt.Printf("loop not detected\n")
}

func bruteforce(head *Node) {
	nodeMap := map[*Node]int {}

	tmp := head

	for tmp != nil {
		if _, ok := nodeMap[tmp]; ok {
			fmt.Printf("loop detected\n")
			return
		}
		nodeMap[tmp] = 1
		tmp = tmp.next
	}

	fmt.Printf("loop not detected\n")
}





func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Printf("%d\n", head.val)
}