package main

import "fmt"


func main() {
	root := buildManualTree()
	ans := preOrder(root)
	fmt.Println(ans)
}

func preOrder(root *TreeNode) []int {
	ans := []int {}
	
	sol(root, &ans)
	return ans
}

func sol(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	sol(root.Left, ans)
	sol(root.Right, ans)
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildManualTree() *TreeNode {
    root := &TreeNode{Val: 4}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 5}

    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = nil

    root.Right.Left = &TreeNode{Val: 7}
    root.Right.Right = &TreeNode{Val: 6}

    return root
}
