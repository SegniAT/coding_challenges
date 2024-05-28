package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var leftCutoff, rightCutoff, dummy *ListNode
	var reverseStart, reverseEnd *ListNode

	dummy = &ListNode{
		Next: head,
	}

	leftCutoff = dummy
	for i := left - 1; i > 0; i-- {
		leftCutoff = leftCutoff.Next
	}
	reverseStart = leftCutoff.Next

	reverseEnd = head
	for iter := right; iter > 1; iter-- {
		reverseEnd = reverseEnd.Next
	}

	rightCutoff = reverseEnd.Next
	reverseEnd.Next = nil

	fmt.Println("leftCutoff: ", leftCutoff)
	fmt.Println("rightCutoff: ", rightCutoff)
	fmt.Println()
	fmt.Println("reverStart: ", reverseStart)
	fmt.Println("reverEnd: ", reverseEnd)

	// reverse
	var prev *ListNode
	curr := reverseStart

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	leftCutoff.Next = reverseEnd
	reverseStart.Next = rightCutoff

	return dummy.Next
}
