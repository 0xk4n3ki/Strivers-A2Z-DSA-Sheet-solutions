package main

import (
	"fmt"

)

type Node struct {
	val int
	next *Node
}

func main() {
	fmt.Printf("-----Case 1-----\n")
	arr1 := []int {2, 4, 3}
	head1 := arr2ll(arr1)
	head1 = reversell(head1)
	printll(head1)

	arr2 := []int {5, 6, 4}
	head2 := arr2ll(arr2)
	head2 = reversell(head2)
	printll(head2)

	head := sumll(head1, head2)
	printll(head)


	fmt.Printf("-----Case 2-----\n")
	arr1 = []int {9, 3, 4}
	head1 = arr2ll(arr1)
	head1 = reversell(head1)
	printll(head1)

	arr2 = []int {7, 4}
	head2 = arr2ll(arr2)
	head2 = reversell(head2)
	printll(head2)

	head = sumll(head1, head2)
	printll(head)
}

func sumll(head1, head2 *Node) *Node {
	d1, d2 := head1, head2
	head := &Node{-1, nil}
	curr := head

	carry := 0

	for d1 != nil || d2 != nil || carry == 1 {
		sum := 0
		if d1 != nil {
			sum += d1.val
			d1 = d1.next
		}

		if d2 != nil {
			sum += d2.val
			d2 = d2.next
		}

		sum += carry
		carry = sum / 10
		tmp := &Node{sum%10, nil}
		curr.next = tmp
		curr = tmp
	}

	
	return reversell(head.next)
}









func reversell(head *Node) *Node {
	curr := head
	var prev *Node = nil

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
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