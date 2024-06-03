package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	list2Head, list2Tail := list2, list2

	for list2Tail.Next != nil {
		list2Tail = list2Tail.Next
	}

	// find ath
	beforeA := list1
	for a > 0 {
		beforeA = beforeA.Next
		a--
	}

	// find b
	afterB := list1
	for b >= 0 {
		afterB = afterB.Next
		b--
	}
	afterB = afterB.Next

	fmt.Println(beforeA)
	fmt.Println(afterB)

	beforeA.Next = list2Head
	list2Tail.Next = afterB

	return list2Head
}
