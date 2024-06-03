package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	stack := [](*ListNode){}

	iter := head

	for iter != nil {
		stack = append(stack, iter)
		iter = iter.Next
	}

	nodeCount := len(stack)
	half := 0
	if nodeCount%2 == 0 {
		half = nodeCount/2 - 1
	} else {
		half = nodeCount / 2
	}

	tailInd := nodeCount - 1

	iter = head
	for half < tailInd {
		next := iter.Next
		tail := stack[tailInd]
		stack[tailInd-1].Next = nil

		iter.Next = tail
		if next != tail {
			tail.Next = next
		}
		iter = next

		tailInd--
	}

}
