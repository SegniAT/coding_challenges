package main

import (
	"math"
	"strconv"
)

func main() {
}

type Stack []int

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	N := len(*s)
	if N == 0 {
		panic("stack.Pop method on empty stack")
	}
	last := (*s)[N-1]
	*s = (*s)[:N-1]
	return last
}

func evalRPN(tokens []string) int {
	stack := Stack{}

	for _, token := range tokens {
		switch token {
		case "+":
			second := stack.Pop()
			first := stack.Pop()
			stack.Push(first + second)
		case "-":
			second := stack.Pop()
			first := stack.Pop()
			stack.Push(first - second)
		case "*":
			second := stack.Pop()
			first := stack.Pop()
			stack.Push(first * second)
		case "/":
			second := stack.Pop()
			first := stack.Pop()
			stack.Push(first / second)
		default:
			num, err := strconv.Atoi(token)
			if err != nil {
				panic("invalid type")
			}
			stack.Push(num)
		}
	}

	return stack.Pop()
}
