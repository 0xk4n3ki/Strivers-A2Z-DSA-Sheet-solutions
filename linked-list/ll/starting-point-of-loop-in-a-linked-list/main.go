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
	start := startingPoint(head)
	fmt.Printf("starting point: %v\n", start)

	start = bruteforce(head)
	fmt.Printf("[bruteforce] starting point: %v\n", start)

	head.next.next.next.next.next.next.next = &Node{372, head.next.next}
	// printll(head)
	start = startingPoint(head)
	fmt.Printf("starting point: %v\n", start)

	start = bruteforce(head)
	fmt.Printf("[bruteforce] starting point: %v\n", start)
}


func startingPoint(head *Node) *Node {
	slow, fast := head, head

	if head == nil || head.next == nil {
		return nil
	}

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if fast == slow {
			slow = head

			for slow != fast {
				slow = slow.next
				fast = fast.next
			}
			return slow
		}
	}

	return nil
}


func bruteforce(head *Node) *Node {
	nodeMap := map[*Node]int {}

	for head != nil {
		if nodeMap[head] != 0 {
			return head
		}
		nodeMap[head] = 1
		head = head.next
	}
	return nil
}


func printll(head *Node) {
	for head.next != nil {
		fmt.Printf("%d->", head.val)
		head = head.next
	}
	fmt.Printf("%d\n", head.val)
}