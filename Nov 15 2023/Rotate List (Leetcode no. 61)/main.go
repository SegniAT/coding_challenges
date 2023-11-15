package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	nodes := 0
	headIter := head
	tail := head

	for headIter != nil {
		nodes++
		if headIter.Next == nil {
			tail = headIter
		}
		headIter = headIter.Next
	}

	rotations := k % nodes
	if rotations == 0 || nodes == 1 {
		return head
	}

	breakoffPt := nodes - rotations
	headIter = head

	for i := 0; i < breakoffPt-1; i++ {
		headIter = headIter.Next
	}

	newHead := headIter.Next
	tail.Next = head
	headIter.Next = nil

	return newHead
}
