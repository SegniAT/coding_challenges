package main

import (
	"errors"
	"fmt"
)

type Item struct {
	index int
	value int
}

var EmptyStackError = errors.New("Stack is empty")

type Stack struct {
	items []Item
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) peek() Item {
	if !s.isEmpty() {
		return (s.items)[len(s.items)-1]
	}
	panic(EmptyStackError)
}

func (s *Stack) push(item Item) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() (Item, error) {
	if !s.isEmpty() {
		lastInd := len(s.items) - 1
		item := s.items[lastInd]
		s.items = s.items[0:lastInd]
		return item, nil
	}

	return Item{}, EmptyStackError
}

// START FROM the END IF YOU SEARCH FOR SOMETHING FORWARD
// START FROM the BEGINNING IF YOU SEARCH FOR SOMETHING BACK
func increasingTriplet(nums []int) bool {
	// Increasing Monotonic Stack
	IMS := newStack()
	prevLess := []int{}
	for range nums {
		prevLess = append(prevLess, 0)
	}

	for i, num := range nums {
		for !IMS.isEmpty() && IMS.peek().value >= num {
			IMS.pop()
		}

		if IMS.isEmpty() {
			prevLess[i] = -1
		} else {
			prevLess[i] = IMS.peek().index
		}

		IMS.push(Item{index: i, value: num})
	}

	//fmt.Println("IMS: ")
	//fmt.Println("nums: ", nums)
	//fmt.Println("IMS: ", IMS)
	//fmt.Println("prev less: ", prevLess)
	//fmt.Println()

	// Decreasing Monotonic Stack
	DMQ := newStack()
	nextGreater := []int{}
	for range nums {
		nextGreater = append(nextGreater, 0)
	}

	for i := len(nums) - 1; i > -1; i-- {
		num := nums[i]
		for !DMQ.isEmpty() && DMQ.peek().value <= num {
			DMQ.pop()
		}

		if DMQ.isEmpty() {
			nextGreater[i] = len(nums)
		} else {
			nextGreater[i] = DMQ.peek().index
		}

		DMQ.push(Item{index: i, value: num})
	}

	//fmt.Println("DMQ: ")
	//fmt.Println("nums: ", nums)
	//fmt.Println("DMQ: ", DMQ)
	//fmt.Println("next greater: ", nextGreater)
	//fmt.Println()

	for i := range nums {
		if prevLess[i] != -1 && nextGreater[i] != len(nums) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 15, 0, 4, 6}))

}
