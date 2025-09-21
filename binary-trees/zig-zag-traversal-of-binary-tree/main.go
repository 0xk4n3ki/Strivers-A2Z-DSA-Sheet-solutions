package main



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int {}
	if root == nil {
		return ans
	}
	
	q := new(Queue)
	q.Push(root)
	flag := 0
	for !q.IsEmpty() {
		level := []int {}
		n := len(q.arr)
		for i := 0; i < n; i++ {
			node := q.Top()
			level = append(level, node.Val)
			q.Pop()

			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		if flag == 1 {
			reverse(level)
		}
		ans = append(ans, level)
		flag ^= 1
	}
	return ans
}

func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

type Queue struct {
	arr []*TreeNode
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
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

func (q *Queue) Top() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	return q.arr[0]
}