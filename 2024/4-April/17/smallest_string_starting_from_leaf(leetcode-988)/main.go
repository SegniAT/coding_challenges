package main

import "fmt"

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	smallestStr := ""
	helper(root, "", &smallestStr)
	return smallestStr
}

// node is never nil
func helper(node *TreeNode, currStr string, smallestStr *string) {
	currStr = fmt.Sprintf("%s%s", numToChar(node.Val), currStr)
	if node.Left == nil && node.Right == nil {
		if *smallestStr == "" || currStr < *smallestStr {
			*smallestStr = currStr
		}
	}

	if node.Left != nil {
		helper(node.Left, currStr, smallestStr)
	}
	if node.Right != nil {
		helper(node.Right, currStr, smallestStr)
	}
}

func numToChar(num int) string {
	return string(rune('a' + num))
}
