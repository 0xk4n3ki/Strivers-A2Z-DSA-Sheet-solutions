package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
		return &TreeNode{val, nil, nil}
	}

	curr := root
	for {
		if curr.Val < val {
			if curr.Right != nil {
				curr = curr.Right
			}else {
				curr.Right = &TreeNode{val, nil, nil}
				break
			}
		}else {
			if curr.Left != nil {
				curr = curr.Left
			}else {
				curr.Left = &TreeNode{val, nil, nil}
				break
			}
		}
	}
	return root
}