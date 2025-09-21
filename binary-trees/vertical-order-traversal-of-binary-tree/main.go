package main

import "sort"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
    ans := [][]int {}
	if root == nil {
		return ans
	}

	nodes := make(map[int]map[int][]int)
	q := new(Queue)
	q.Push(Item{root, 0, 0})
	
	for !q.IsEmpty() {
		curr := q.Top()
		q.Pop()

		node := curr.Node
		x, y := curr.x, curr.y

		if _, ok := nodes[x]; !ok {
			nodes[x] = make(map[int][]int)
		}
		nodes[x][y] = append(nodes[x][y], node.Val)

		if node.Left != nil {
			q.Push(Item{node.Left, x-1, y+1})
		}
		if node.Right != nil {
			q.Push(Item{node.Right, x+1, y+1})
		}
	}

	xKeys := make([]int, 0, len(nodes))
	for k := range nodes {
		xKeys = append(xKeys, k)
	}
	sort.Ints(xKeys)

	for _, x := range xKeys {
		col := []int {}

		yKeys := []int {}
		for y := range nodes[x] {
			yKeys = append(yKeys, y)
		}
		sort.Ints(yKeys)

		for _, y := range yKeys {
			vals := nodes[x][y]
			sort.Ints(vals)
			col = append(col, vals...)
		}
		ans = append(ans, col)
	}
	return ans
}

type Item struct {
	Node *TreeNode
	x int
	y int
}

type Queue struct {
	arr []Item
}

func (q *Queue) IsEmpty() bool{
	return len(q.arr) == 0
}

func (q *Queue) Push(ele Item) {
	q.arr = append(q.arr, ele)
}

func (q *Queue) Top() Item {
	if q.IsEmpty() {
		return Item{}
	}
	return q.arr[0]
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[1:]
}