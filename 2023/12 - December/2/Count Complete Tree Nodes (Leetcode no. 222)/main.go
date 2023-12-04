package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// get height
	level, levelITer := 0, root
	for root.Left != nil {
		level++
		levelITer = levelITer.Left
	}

	totalNodes := int(math.Pow(2, float64(level+1)) - 1)

	return totalNodes
}
