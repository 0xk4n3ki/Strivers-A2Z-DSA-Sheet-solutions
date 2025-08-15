package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {1, 0, 2, 0, 1}
	head := arr2ll(arr)
	printll(head)

	head = sortll(head)
	
	printll(head)
}


func sortll(head *Node) *Node {
	zeroHead, onesHead, twosHead := &Node{-1, nil}, &Node{-1, nil}, &Node{-1, nil}

	zero, one, two := zeroHead, onesHead, twosHead
	curr := head

	for curr != nil {
		ele := curr.val

		if ele == 0 {
			zero.next = curr
			zero = zero.next
		}else if ele == 1 {
			one.next = curr
			one = one.next
		}else {
			two.next = curr
			two = two.next
		}

		curr = curr.next
	}

	if onesHead.next != nil {
		zero.next = onesHead.next
	}
	one.next = twosHead.next
	two.next = nil

	return zeroHead.next
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