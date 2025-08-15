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

	head = bruteforceApp(head)

	printdll(head)

	head = optimal(head)

	printdll(head)
}

func optimal(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	for head.next != nil {
		head.next, head.back = head.back, head.next
		head = head.back
	}

	head.next, head.back = head.back, head.next

	return head
}

func bruteforceApp(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	cur := &Node{head.val, nil, nil}

	head = head.next

	for head != nil {
		tmp := &Node{head.val, cur, nil}
		cur.back = tmp

		cur = tmp
		head = head.next
	}

	return cur

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