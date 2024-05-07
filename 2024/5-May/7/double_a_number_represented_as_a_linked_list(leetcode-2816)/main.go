package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	carryOver := helper(head)
	curr := head
	for carryOver > 0 {
		num := carryOver
		carryOver = num / 10
		curr = &ListNode{
			Val:  num % 10,
			Next: curr,
		}
	}

	return curr
}

func helper(head *ListNode) int {
	if head.Next == nil {
		num := head.Val * 2
		head.Val = num % 10
		return num / 10
	}

	carryOver := helper(head.Next)
	num := head.Val*2 + carryOver

	head.Val = num % 10
	return num / 10
}
