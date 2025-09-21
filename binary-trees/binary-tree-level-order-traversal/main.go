package main


//  Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func levelOrder(root *TreeNode) [][]int {
    ans := [][]int {}
	if root == nil {
		return ans
	}

	q := new(Queue)
	q.Push(root)

	for !q.IsEmpty() {
		n := len(q.arr)
		level := []int {}
		for i := 0; i < n; i++ {
			node := q.Front()
			q.Pop()
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			level = append(level, node.Val)
		}
		ans = append(ans, level)
	}
	return ans
}

type Queue struct {
	arr []*TreeNode
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Front() *TreeNode {
	return q.arr[0]
}

func (q *Queue) Push(node *TreeNode) {
	q.arr = append(q.arr, node)
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[1:]
}