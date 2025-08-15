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

	head = bruteforce(head, 3)
	printll(head)

	head = bruteforce(head, 4)
	printll(head)


	arr = []int {1, 2, 3, 4, 5, 372, 12, 27}
	head = arr2ll(arr)
	printll(head)

	head = optimal(head, 8)
	printll(head)

	head = optimal(head, 3)
	printll(head)
}

func optimal(head *Node, i int) *Node {
	slow, fast := head, head

	for j := i; j > 0; j-- {
		fast = fast.next
	}

	if fast == nil {
		newHead := head.next
		head.next = nil
		return newHead
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	delNode := slow.next
	slow.next = delNode.next
	delNode.next = nil

	return head
}


func bruteforce(head *Node, i int) *Node {
	n := lenll(head)

	if i > n {
		return head
	}

	if n == i {
		newHead := head.next
		head.next = nil
		return newHead
	}

	j := n - i
	curr := head
	
	for curr != nil {
		j--

		if j == 0 {
			delNode := curr.next
			curr.next = curr.next.next
			delNode.next = nil

			return head
		}
		curr = curr.next
	}

	return nil
}


func lenll(head *Node) int {
	n := 0

	curr := head
	for curr != nil {
		n++
		curr = curr.next
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