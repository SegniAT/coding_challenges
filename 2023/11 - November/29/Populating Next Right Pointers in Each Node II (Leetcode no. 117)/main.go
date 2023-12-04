package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	nextChild := findNextChild(root.Next)

	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}

	if root.Right != nil {
		root.Right.Next = nextChild
	} else if root.Left != nil {
		root.Left.Next = nextChild
	}

	// ****************************
	// DO RIGHT FIRST SINCE WE HAVE TO FINISH THE LINKS(LEFT -> RIGHT) FROM THE RIGHT FIRST ...
	// TO ITERATE OVER THE ROOT'S NEXT FULLY WITH findNextChild
	connect(root.Right)
	connect(root.Left)

	return root
}

func findNextChild(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.Left != nil {
		return node.Left
	}

	if node.Right != nil {
		return node.Right
	}
	return findNextChild(node.Next)
}
