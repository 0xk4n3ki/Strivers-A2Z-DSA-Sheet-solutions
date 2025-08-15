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
	arr := []int {1, 1, 3, 3, 4, 5}
	head := arr2dll(arr)
	printdll(head)
	head = removeDup(head)
	printdll(head)


	arr = []int {1, 1, 1, 1, 1, 2}
	head = arr2dll(arr)
	printdll(head)
	head = removeDup(head)
	printdll(head)
}

func removeDup(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	
	curr := head
	last := curr.val
	curr = curr.next

	for curr != nil {
		if curr.val == last {
			prev := curr.back
			front := curr.next

			if front != nil {
				front.back = prev
			}
			prev.next = front
			curr = front
		}else {
			last = curr.val
			curr = curr.next
		}
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