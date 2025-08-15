package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
	back *Node
}

func main() {
	arr := []int {1, 2, 3, 1, 4}
	head := arr2dll(arr)
	target := 1
	printdll(head)
	head = deleteAll(head, target)
	printdll(head)

	arr = []int {2, 3, -1, 4, 2}
	head = arr2dll(arr)
	target = 2
	printdll(head)

	head = deleteAll(head, target)
	printdll(head)

	arr = []int {2, 2, 2, 2, 2, 2}
	head = arr2dll(arr)
	target = 2
	printdll(head)

	head = deleteAll(head, target)
	printdll(head)
}

func deleteAll(head *Node, target int) *Node {
	curr := head

	for curr != nil && curr.val == target {
		next := curr.next
		curr.next = nil
		if next != nil {
			next.back = nil
		}
		curr = next
		head = curr
	}

	for curr != nil {
		if curr.val == target {
			prev := curr.back
			next := curr.next

			if prev != nil {
				prev.next = next
			}
			if next != nil {
				next.back = prev
			}
		}
		curr = curr.next
	}
	return head
}










func arr2dll(arr []int) *Node {
	head := &Node{arr[0], nil, nil}

	prev := head

	for i := 1; i < len(arr); i++ {
		temp := &Node{arr[i], nil, prev}

		prev.next = temp
		prev = temp
	}

	return head
}


func printdll(head *Node) {
	for head != nil {
		fmt.Printf("%d", head.val)

		if head.next != nil {
			fmt.Printf(" <-> ")
		}
		head = head.next
	}
	fmt.Println()
}