package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
    ans := []int {}
    if root == nil {
		return ans
	}

	st := new(Stack)
	st.Push(root)
	for !st.IsEmpty() {
		root = st.Top()
		st.Pop()
		ans = append(ans, root.Val)
		if root.Right != nil {
			st.Push(root.Right)
		}
		if root.Left != nil {
			st.Push(root.Left)
		}
	}
	return ans
}

type Stack struct {
	arr []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.arr = append(s.arr, node)
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
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