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
	mid := bruteForce(head)
	fmt.Printf("mid: %d\n", mid)
	mid = optimal(head)
	fmt.Printf("mid: %d\n", mid)

	arr = []int {1, 2, 3, 4, 5, 6}
	head = arr2ll(arr)
	printll(head)
	mid = bruteForce(head)
	fmt.Printf("mid: %d\n", mid)
	mid = optimal(head)
	fmt.Printf("mid: %d\n", mid)
}


func optimal(head *Node) int {
	fast, slow := head, head

	for fast != nil && fast.next != nil  {
		slow = slow.next

		fast = fast.next.next
	}

	return slow.val
}


func bruteForce(head *Node) int {
	n := lenll(head)
	m := n/2 + 1


	n = 1

	for n < m {
		head = head.next
		n++
	}

	return head.val
}

func lenll(head *Node) int {
	n := 0

	for head != nil {
		n++
		head = head.next
	}
	return n
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