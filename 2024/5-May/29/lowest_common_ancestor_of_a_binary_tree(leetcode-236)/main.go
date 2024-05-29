package main

import "fmt"

func main() {
	left := &TreeNode{
		Val: 2,
	}

	right := &TreeNode{
		Val: 3,
	}

	root := &TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}

	fmt.Println(lowestCommonAncestor(root, left, right))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
