package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	iter := head
	for iter != nil {
		next := iter.Next
		iter.Next = prev
		prev = iter
		iter = next
	}

	return prev
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	nodes := 0
	iter := head

	for iter != nil {
		iter = iter.Next
		nodes++
	}

	middle := nodes / 2
	iter = head
	for middle > 0 {
		iter = iter.Next
		middle--
	}

	var reversedHalf *ListNode
	if nodes%2 == 0 {
		reversedHalf = reverseList(iter)
	} else {
		reversedHalf = reverseList(iter.Next)
	}

	iter1 := head
	iter2 := reversedHalf

	for iter2 != nil {
		if iter1.Val != iter2.Val {
			return false
		}
		iter1 = iter1.Next
		iter2 = iter2.Next
	}

	return true
}
