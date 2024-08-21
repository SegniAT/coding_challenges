package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
}

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("popping from an empty stack")
	}

	N := s.Size()
	return (*s)[N-1]
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("popping from an empty stack")
	}

	N := s.Size()
	last := (*s)[N-1]
	(*s) = (*s)[:N-1]

	return last
}

func isValid(s string) bool {
	stack := NewStack[rune]()

	opp := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack.Push(char)
		} else {
			if stack.Peek() != opp[char] {
				return false
			}

			stack.Pop()
		}
	}

	return stack.IsEmpty()
}
