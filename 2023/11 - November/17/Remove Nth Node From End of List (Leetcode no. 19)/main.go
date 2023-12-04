package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Two pass
func removeNthFromEndTwoPass(head *ListNode, n int) *ListNode {
	ans := &ListNode{
		Next: head,
	}

	size := 0
	iter := ans
	for iter.Next != nil {
		size++
		iter = iter.Next
	}

	steps := size - n

	iter = ans

	for i := 0; i < steps; i++ {
		iter = iter.Next
	}

	iter.Next = iter.Next.Next

	return ans.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ans := &ListNode{
		Next: head,
	}

	// create a left and right pointers
	left, right := ans, ans

	// create the window
	for i := 0; i < n; i++ {
		right = right.Next
	}

	// move the window until the right's next is nil
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	// remove the element
	left.Next = left.Next.Next

	return ans.Next
}
