package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive solution
func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// iterative solution
func reverseList(head *ListNode) *ListNode {
	iter := head

	var prev *ListNode
	for iter != nil {
		next := iter.Next
		iter.Next = prev
		prev = iter
		iter = next
	}

	return prev
}
