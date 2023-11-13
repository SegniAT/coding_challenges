package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Naive
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	l1Iter, l2Iter, resIter := l1, l2, res

	remainder := 0
	for l1Iter != nil || l2Iter != nil || remainder > 0 {
		resIter.Next = &ListNode{}
		resIter = resIter.Next

		l1IterVal := 0
		if l1Iter != nil {
			l1IterVal = l1Iter.Val
			l1Iter = l1Iter.Next
		}

		l2IterVal := 0
		if l2Iter != nil {
			l2IterVal = l2Iter.Val
			l2Iter = l2Iter.Next
		}

		currSum := l1IterVal + l2IterVal + remainder
		resIter.Val = currSum % 10
		remainder = currSum / 10
	}

	return res.Next
}
