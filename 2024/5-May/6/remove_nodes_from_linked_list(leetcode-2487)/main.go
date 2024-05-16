package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	stack := []*ListNode{}

	iter := head

	for iter != nil {
		for len(stack) > 0 && stack[len(stack)-1].Val < iter.Val {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, iter)
		iter = iter.Next
	}

	for i := 0; i < len(stack)-1; i++ {
		stack[i].Next = stack[i+1]
	}

	return stack[0]
}
