package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {1, 2, 3, 4, 5}
	head := arr2ll(arr)
	printll(head)

	head = deleteMid(head)
	printll(head)

	arr = []int {1, 2, 3, 4}
	head = arr2ll(arr)
	printll(head)

	head = deleteMid(head)
	printll(head)
}


func deleteMid(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	
	slow, fast := head, head

	var prev *Node = nil

	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	prev.next = slow.next
	slow.next = nil

	return head
}







func arr2ll(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{arr[0], nil}
	cur := head

	for i := 1; i < len(arr); i++ {
		tmp := &Node{arr[i], nil}
		cur.next = tmp

		cur = tmp
	}

	return head
}

func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Printf("%d\n", head.val)
}