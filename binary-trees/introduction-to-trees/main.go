package main

import "fmt"

func main() {
	root := &Node{1, nil, nil}
	root.left = &Node{2, nil, nil}
	root.right = &Node{3, nil, nil}
	root.left.left = new(Node)

	fmt.Println(root)
	fmt.Println(root.left)
	fmt.Println(root.right)
	fmt.Println(root.left.left)
}

type Node struct {
	val int
	left *Node
	right *Node
}