package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CeilNode(root *TreeNode, key int) int {
	ceil := -1

	for root != nil {
		if root.Val == key {
			ceil = key
			break
		}else if root.Val > key {
			ceil = root.Val
			root = root.Left
		}else {
			root = root.Right
		}
	}
	return ceil
}