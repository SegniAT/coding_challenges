package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root)
}

func helper(node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		return node
	}

	rightTemp := node.Right
	leftHead := node.Left
	var leftLeaf *TreeNode
	if node.Left != nil {
		leftLeaf = helper(node.Left)
	}

	node.Left = nil

	if leftHead != nil {
		node.Right = leftHead
	}

	if leftLeaf != nil {
		leftLeaf.Right = rightTemp
	}

	right := helper(node.Right)

	return right
}
