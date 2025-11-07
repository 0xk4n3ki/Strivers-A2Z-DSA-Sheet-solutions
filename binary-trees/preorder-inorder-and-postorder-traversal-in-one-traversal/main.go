package main

import (
	"fmt"
)

func main() {
	arr := []int {1, 2, 3, 4, 5, 6, 7, -1, -1, 8, -1, -1, -1, 9, 10}
	root := &Node{arr[0], nil, nil}
	arr2tree(arr, root, 0)

	pre, in, post := traversal(root)
	fmt.Printf("pre: %v\n", pre)
	fmt.Printf("in: %v\n", in)
	fmt.Printf("post: %v\n", post)
}

type pair struct {
	node *Node 
	state int
}

func traversal(root *Node) ([]int, []int, []int) {
	pre, in, post := []int{}, []int{}, []int{}
	st := []pair {}

	if root == nil {
		return pre, in, post
	}
	st = append(st, pair{root, 1})

	for len(st) > 0 {
		curr := st[len(st)-1]
		st = st[:len(st)-1]

		switch curr.state {
		case 1:
			pre = append(pre, curr.node.val)
			curr.state = 2
			st = append(st, curr)
			if curr.node.left != nil {
				st = append(st, pair{curr.node.left, 1})
			}
		case 2:
			in = append(in, curr.node.val)
			curr.state = 3
			st = append(st, curr)
			if curr.node.right != nil {
				st = append(st, pair{curr.node.right, 1})
			}
		case 3:
			post = append(post, curr.node.val)			
		}
	}
	return pre, in, post
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