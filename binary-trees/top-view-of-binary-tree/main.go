package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func TopView(root *TreeNode) []int {
	ans := []int {}
	if root == nil {
		return ans
	}

	nodes := map[int]*TreeNode {}
	q := new(Queue)
	q.Push(Item{root, 0})

	for !q.IsEmpty() {
		curr := q.Top()
		q.Pop()
		node := curr.Node
		col := curr.col
		
		if _, ok := nodes[col]; !ok {
			nodes[col] = node
		}
		if node.Left != nil {
			q.Push(Item{node.Left, col-1})
		}
		if node.Right != nil {
			q.Push(Item{node.Right, col+1})
		}
	}

	cols := []int {}
	for col := range nodes {
		cols = append(cols, col)
	}
	sort.Ints(cols)

	for _, col := range cols {
		ans = append(ans, nodes[col].Val)
	}
	return ans
}

type Item struct {
	Node *TreeNode
	col int
}

type Queue struct {
	arr []Item
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Push(item Item) {
	q.arr = append(q.arr, item)
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