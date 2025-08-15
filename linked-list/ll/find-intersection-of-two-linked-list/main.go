package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	arr1 := []int {1, 3, 1, 2, 4, 5, 17, 19}
	arr2 := []int {3, 2, 4}
	head1 := arr2ll(arr1)
	head2 := arr2ll(arr2)

	printll(head1)
	printll(head2)

	usingOneMap(head1, head2)
	diffInLen(head1, head2)
	optimal(head1, head2)

	head2 = &Node{372, nil}
	head2.next = &Node{373, nil}
	head2.next.next = head1.next.next.next.next

	printll(head2)
	usingOneMap(head1, head2)
	diffInLen(head1, head2)
	optimal(head1, head2)
}

func optimal(head1, head2 *Node) {
	d1, d2 := head1, head2

	for d1 != d2 {
		if d1 == nil {
			d1 = head2
		}else {
			d1 = d1.next
		}

		if d2 == nil {
			d2 = head1
		}else {
			d2 = d2.next
		}
	}

	if d1 != nil {
		fmt.Printf("[optimal] ans: %d\n", d1.val)
		return
	}
	fmt.Printf("[optimal] ans: %d\n", -1)
}

func diffInLen(head1, head2 *Node) {
	n1, n2 := lenll(head1), lenll(head2)

	var toInc, nor *Node
	var diff int

	if n1 > n2 {
		toInc = head1
		nor = head2
		diff = n1 - n2
	}else {
		toInc = head2
		nor = head1
		diff = n2 - n1
	}

	for diff > 0 {
		diff--
		toInc = toInc.next
	}

	for toInc != nil && nor != nil {
		if toInc == nor {
			fmt.Printf("[diffInLen] ans: %d\n", toInc.val)
			return
		}
		toInc = toInc.next
		nor = nor.next
	}
	fmt.Printf("[diffInLen] ans: %d\n", -1)
}

func lenll(head *Node) int {
	curr := head
	n := 0
	for curr != nil {
		n++
		curr = curr.next
	}
	return n
}

func usingOneMap(head1, head2 *Node) {
	nodeMap := map[*Node]struct{} {}

	cur := head1
	for cur != nil {
		nodeMap[cur] = struct{}{}
		cur = cur.next
	}

	cur = head2
	for cur != nil {
		if _, ok := nodeMap[cur]; ok {
			fmt.Printf("[usingOneMap] ans: %d\n", cur.val)
			return
		}
		cur = cur.next
	}
	fmt.Printf("[usingOneMap] ans: %d\n", -1)
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