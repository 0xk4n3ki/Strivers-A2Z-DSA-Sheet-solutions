package main

import "fmt"

type Node struct {
	val int
	next *Node
	back *Node
}

func main() {
	arr := []int {2, 4, 5, 7, 9, 12, 45}

	head := arr2dll(arr)

	printdll(head)

	head = insertEnd(head, 6)

	printdll(head)
}

func insertEnd(head *Node, n int) *Node {
	if head == nil {
		return &Node{
			val: n,
			next: nil,
			back: nil,
		}
	}

	cur := head

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &Node{n, nil, cur}

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