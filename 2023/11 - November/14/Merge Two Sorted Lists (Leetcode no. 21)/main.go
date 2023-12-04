package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{}
	ansIter, l1Iter, l2Iter := ans, list1, list2

	for l1Iter != nil || l2Iter != nil {
		if l1Iter == nil {
			ansIter.Next = l2Iter
			break
		}

		if l2Iter == nil {
			ansIter.Next = l1Iter
			break
		}

		if l1Iter.Val <= l2Iter.Val {
			ansIter.Next = l1Iter
			l1Iter = l1Iter.Next
		} else {
			ansIter.Next = l2Iter
			l2Iter = l2Iter.Next
		}

		ansIter = ansIter.Next
	}

	return ans.Next
}
