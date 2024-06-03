package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return helper(root, false)
}

func helper(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}

	if isLeft && node.Left == nil && node.Right == nil {
		return node.Val
	}

	left, right := 0, 0

	left = helper(node.Left, true)
	right = helper(node.Right, false)

	return left + right
}
