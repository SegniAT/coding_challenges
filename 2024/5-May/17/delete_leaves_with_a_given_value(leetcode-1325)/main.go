package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	if isLeaf(root) {
		if root.Val == target {
			return nil
		} else {
			return root
		}
	}

	left := removeLeafNodes(root.Left, target)
	right := removeLeafNodes(root.Right, target)

	root.Left = left
	root.Right = right

	if left == nil && right == nil && root.Val == target {
		return nil
	}

	return root
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
