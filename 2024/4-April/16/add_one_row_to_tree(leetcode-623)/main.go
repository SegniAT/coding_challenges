package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}

	helper(root, val, depth, 1)
	return root
}

func helper(node *TreeNode, val, depth, currDepth int) {
	if node == nil {
		return
	}

	if currDepth == depth-1 {
		left := node.Left
		right := node.Right

		node.Left = &TreeNode{
			Val:  val,
			Left: left,
		}

		node.Right = &TreeNode{
			Val:   val,
			Right: right,
		}
	}

	helper(node.Left, val, depth, currDepth+1)
	helper(node.Right, val, depth, currDepth+1)
}
