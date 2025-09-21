package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    ans := []int {}
	st := new(Stack)
	node := root

	for {
		if node != nil {
			st.Push(node)
			node = node.Left
		}else {
			if st.IsEmpty() {
				break
			}
			node = st.Top()
			st.Pop()
			ans = append(ans, node.Val)
			node = node.Right
		}
	}
	return ans
}

type Stack struct {
	arr []*TreeNode
}

func (st *Stack) IsEmpty() bool {
	return len(st.arr) == 0
}

func (st *Stack) Push(val *TreeNode) {
	st.arr = append(st.arr, val)
}

func (st *Stack) Top() *TreeNode {
	if st.IsEmpty() {
		return nil
	}
	return st.arr[len(st.arr)-1]
}

func (st *Stack) Pop() {
	if st.IsEmpty() {
		return
	}
	st.arr = st.arr[:len(st.arr)-1]
}