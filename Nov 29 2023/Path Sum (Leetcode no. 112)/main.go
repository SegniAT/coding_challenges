package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: nil,
	}

	fmt.Println(hasPathSum(root, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return helper(root, targetSum)
}

func helper(root *TreeNode, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}

	left := false
	if root.Left != nil {
		left = helper(root.Left, targetSum-root.Val)
	}

	right := false
	if root.Right != nil {
		right = helper(root.Right, targetSum-root.Val)
	}

	return left || right
}
