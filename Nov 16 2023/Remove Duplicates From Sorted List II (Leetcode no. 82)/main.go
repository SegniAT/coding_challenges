package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	ans := &ListNode{}
	if head == nil {
		return head
	}

	iter := head
	ansIter := ans
	var prev, next *ListNode = nil, nil

	for iter != nil {
		next = iter.Next

		if ((prev != nil && prev.Val < iter.Val) || prev == nil) && ((next != nil && iter.Val < next.Val) || next == nil) {
			ansIter.Next = iter
			ansIter = ansIter.Next
		}

		prev = iter
		iter = iter.Next
	}

	ansIter.Next = nil

	return ans.Next
}
