package main

import (
	"fmt"
)

func main() {
	arr := []int {1, 2, 5, -1, -1, 4, 6, -1, -1, -1, -1, 5}
	root := &Node{arr[0], nil, nil}
	arr2tree(arr, root, 0)

	ans := levelTraverse(root)
	fmt.Printf("height: %d\n", ans)
	fmt.Printf("height: %d\n", height(root))
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	lh := height(root.left)
	rh := height(root.right)

	if lh > rh {
		return lh+1
	}
	return rh+1
}


func levelTraverse(root *Node) int {
	if root == nil {
		return 0
	}
	queue := []*Node {root}
	level := 0

	for len(queue) != 0 {
		level++

		num := len(queue)
		for i := 0; i < num; i++ {
			node := queue[0]
			queue = queue[1:]

			fmt.Printf("%d ", node.val)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	fmt.Println()
	return level
}

type Node struct {
	val int
	left *Node
	right *Node
}

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

func arr2tree(arr []int, root *Node, i int) {
	n := len(arr)

	if left(i) < n && arr[left(i)] != -1 {
		root.left = &Node{arr[left(i)], nil, nil}
		arr2tree(arr, root.left, left(i))
	}

	if right(i) < n && arr[right(i)] != -1 {
		root.right = &Node{arr[right(i)], nil, nil}
		arr2tree(arr, root.right, right(i))
	}
}