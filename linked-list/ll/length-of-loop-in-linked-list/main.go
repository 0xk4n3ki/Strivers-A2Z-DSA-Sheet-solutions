package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	head := &Node{1, nil}
	head.next = &Node{2, nil}
	head.next.next = &Node{3, nil}
	head.next.next.next = &Node{4, nil}
	head.next.next.next.next = &Node{9, nil}
	head.next.next.next.next.next = &Node{9, nil}
	head.next.next.next.next.next.next = &Node{5, nil}

	printll(head)
	n := bruteforce(head)
	fmt.Printf("len of loop: %d\n", n)

	n = optimal(head)
	fmt.Printf("[optimal] len of loop: %d\n", n)

	head.next.next.next.next.next.next.next = &Node{372, head.next.next}
	n = bruteforce(head)
	fmt.Printf("len of loop: %d\n", n)

	n = optimal(head)
	fmt.Printf("[optimal] len of loop: %d\n", n)
}


func optimal(head *Node) int {
	if head == nil || head.next == nil {
		return 0
	}

	slow, fast := head, head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			count := 1

			fast = fast.next
			for fast != slow {
				count++
				fast = fast.next
			}
			return count
		}
	}
	return 0
}


func bruteforce(head *Node) int {
	nodeMap := map[*Node]int {}

	index := 0
	for head != nil {
		if i, ok := nodeMap[head]; ok {
			// fmt.Printf("map: %v\n", nodeMap)
			return index - i + 1
		}

		index++
		nodeMap[head] = index

		head = head.next
	}
	return 0
}


func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Printf("%d\n", head.val)
}