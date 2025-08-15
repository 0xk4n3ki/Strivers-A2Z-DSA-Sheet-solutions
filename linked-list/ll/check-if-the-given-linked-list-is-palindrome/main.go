package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {1, 2, 3, 2, 1}
	head := arr2ll(arr)

	printll(head)

	ans := checkPalindrome(head)
	fmt.Printf("ans: %t\n", ans)

	ans = optimal(head)
	fmt.Printf("[optimal] ans: %t\n", ans)


	arr = []int {1, 2, 3, 2, 3}
	head = arr2ll(arr)
	printll(head)
	ans = checkPalindrome(head)
	fmt.Printf("ans: %t\n", ans)

	ans = optimal(head)
	fmt.Printf("[optimal] ans: %t\n", ans)
}


func optimal(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	slow, fast := head, head

	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	secondHalf := reversell(slow.next)

	first, second := head, secondHalf

	for second != nil {
		if first.val != second.val {
			// reversell(secondHalf)
			return false
		}

		first, second = first.next, second.next
	}

	// reversell(secondHalf)
	return true
}

func reversell(head *Node) *Node {
	tmp := head
	var prev *Node = nil

	for tmp != nil {
		front := tmp.next

		tmp.next = prev

		prev = tmp

		tmp = front
	}
	return prev
}


func checkPalindrome(head *Node) bool {
	temp := head
	arr := []int {}

	for temp != nil {
		arr = append(arr, temp.val)

		temp = temp.next
	}

	temp = head
	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] != temp.val {
			return false
		}
		temp = temp.next
	}
	return true
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