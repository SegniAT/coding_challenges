package main

func main() {
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return getBool(root.Val)
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	res := false
	if root.Val == 2 {
		res = left || right
	}else if root.Val == 3 {
		res = left && right
	}

	return res
}

func getBool(val int) bool {
	return val == 1
}
