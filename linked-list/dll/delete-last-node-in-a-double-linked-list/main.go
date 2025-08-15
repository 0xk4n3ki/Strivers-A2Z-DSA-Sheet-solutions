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
	arr := []int {2, 4, 5, 7, 9, 12, 45}

	head := arr2dll(arr)

	printdll(head)

	head =	deleteLast(head)

	printdll(head)
}

func deleteLast(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	cur := head
	for cur.next.next != nil {
		cur = cur.next
	}

	cur.next = nil

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