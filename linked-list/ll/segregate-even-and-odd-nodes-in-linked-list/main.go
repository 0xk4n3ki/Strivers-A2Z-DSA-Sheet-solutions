package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	arr := []int {1, 2, 3, 4, 5, 6}
	head := arr2ll(arr)
	printll(head)

	head = segregatell(head)
	printll(head)
}

func segregatell(head *Node) *Node {
	oddHead, eveHead := &Node{-1, nil}, &Node{-1, nil}
	oddTail, evenTail := oddHead, eveHead

	curr := head

	for curr != nil {
		tmp := curr
		curr = curr.next

		tmp.next = nil

		ele := tmp.val

		if ele % 2 == 0 {
			evenTail.next = tmp
			evenTail = tmp
		}else {
			oddTail.next = tmp
			oddTail = tmp
		}
	}

	evenTail.next = oddHead.next

	return eveHead.next
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