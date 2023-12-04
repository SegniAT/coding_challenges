package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 0)
}

func helper(node *TreeNode, curr int) int {
	newCurr := curr*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return newCurr
	}

	left := 0
	if node.Left != nil {
		left = helper(node.Left, newCurr)
	}

	right := 0
	if node.Right != nil {
		right = helper(node.Right, newCurr)
	}

	return left + right
}
