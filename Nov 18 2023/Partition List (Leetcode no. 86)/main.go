package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lowerNodes := &ListNode{}
	eqOrGreaterNodes := &ListNode{}

	mainIter := head
	lowerNodesIter := lowerNodes
	eqOrGreaterNodesIter := eqOrGreaterNodes

	for mainIter != nil {
		if mainIter.Val < x {
			lowerNodesIter.Next = mainIter
			lowerNodesIter = lowerNodesIter.Next
		} else {
			eqOrGreaterNodesIter.Next = mainIter
			eqOrGreaterNodesIter = eqOrGreaterNodesIter.Next
		}
		mainIter = mainIter.Next
	}

	eqOrGreaterNodesIter.Next = nil
	lowerNodesIter.Next = eqOrGreaterNodes.Next

	return lowerNodes.Next
}
