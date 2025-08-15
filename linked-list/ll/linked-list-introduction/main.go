package main

import "fmt"

type Node struct {
	val int
	next *Node
}

func newNode(n int) *Node {
	return &Node{
		val: n,
	}
}

func main() {
	arr := []int {2, 3, 4, 5}

	y := newNode(arr[0])
	fmt.Printf("y: %v\n", y)

	fmt.Printf("y: %p\n", y)
	fmt.Println(y.val)
	fmt.Println(y.next)
}