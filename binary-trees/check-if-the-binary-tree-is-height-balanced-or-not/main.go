package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {3, 9, 20, -1, -1, 15, 7}
	root := &Node{arr[0], nil, nil}
	arr2tree(arr, root, 0)

	fmt.Printf("ans: %t\n", brute(root))


	arr = []int {1, 3, 2, 5, 4, -1, -1, 7, 6}
	root = &Node{arr[0], nil, nil}
	arr2tree(arr, root, 0)

	fmt.Printf("ans: %t\n", brute(root))
}

func brute(root *Node) bool {
	if root == nil {
		return true
	}
	lh := getHeight(root.left)
	rh := getHeight(root.right)

	if math.Abs(float64(lh-rh)) <= 1 && brute(root.left) && brute(root.right) {
		return true
	}
	return false
}

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}
	lh := getHeight(root.left)
	rh := getHeight(root.right)

	if lh > rh {
		return lh+1
	}
	return rh+1
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