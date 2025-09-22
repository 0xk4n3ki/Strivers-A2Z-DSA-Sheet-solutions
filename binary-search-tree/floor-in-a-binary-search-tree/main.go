package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func floorBST(root *TreeNode, key int) int {
	floor := -1
	for root != nil {
		if root.Val == key {
			floor = key
			break
		}else if root.Val > key {
			root = root.Left
		}else {
			floor = root.Val
			root = root.Right
		}
	}
	return floor
}