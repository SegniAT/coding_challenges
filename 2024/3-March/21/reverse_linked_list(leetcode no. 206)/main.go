package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

// iterative
func reverseListOne(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	nodes := []*ListNode{}

	iter := head

	for iter != nil {
		nodes = append(nodes, iter)
		iter = iter.Next
	}

	nodesLen := len(nodes)

	for i := 0; i < nodesLen; i++ {
		if i != 0 {
			nodes[i].Next = nodes[i-1]
		} else {
			nodes[i].Next = nil
		}
	}

	return nodes[nodesLen-1]
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	tail := reverseList(head.Next)
	curr := head
	curr.Next.Next = curr
	curr.Next = nil

	return tail
}
