package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {3, 4, 2, 1, 5}
	head := arr2ll(arr)
	printll(head)

	head = inplaceSorting(head)
	printll(head)
}


func inplaceSorting(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	secHalf := slow.next
	slow.next = nil
	firstHalf := head

	firstHalf = inplaceSorting(firstHalf)
	secHalf = inplaceSorting(secHalf)

	return mergeTwoll(firstHalf, secHalf)
}

func mergeTwoll(left, right *Node) *Node {
	dummyNode := &Node{-1, nil}
	temp := dummyNode

	for left != nil && right != nil {
		if left.val <= right.val {
			temp.next = left
			left = left.next
		}else {
			temp.next = right
			right = right.next
		}
		temp = temp.next
	}

	if left != nil {
		temp.next = left
	}else {
		temp.next = right
	}

	return dummyNode.next
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