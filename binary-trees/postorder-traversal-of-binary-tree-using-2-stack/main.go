package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
    st1, st2 := new(Stack), new(Stack)
	st1.Push(root)

	for {
		node := st1.Top()
		if node == nil {
			break
		}
		st2.Push(node)
		st1.Pop()

		if node.Left != nil {
			st1.Push(node.Left)
		}
		if node.Right != nil {
			st1.Push(node.Right)
		}
	}

	ans := []int {}
	for !st2.IsEmpty() {
		ans = append(ans, st2.Top().Val)
		st2.Pop()
	}
	return ans
}

type Stack struct {
	arr []*TreeNode
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Push(node *TreeNode) {
	s.arr = append(s.arr, node)
}

func (s *Stack) Top() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.arr = s.arr[:len(s.arr)-1]
}