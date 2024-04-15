package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(node *TreeNode, num int) int {
	if node == nil {
		return num
	}

	if node.Left == nil && node.Right == nil {
		return (num * 10) + node.Val
	}

	left, right := 0, 0

	if node.Left != nil {
		left = helper(node.Left, (num*10)+node.Val)
	}
	if node.Right != nil {
		right = helper(node.Right, (num*10)+node.Val)
	}

	return left + right
}
