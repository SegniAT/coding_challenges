package main

import "container/list"

func main() {

}

type Stack[T any] struct {
	list *list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: list.New(),
	}
}

func (s *Stack[T]) Push(value T) {
	s.list.PushBack(value)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	e := s.list.Back()
	s.list.Remove(e)
	return e.Value.(T)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.Len() == 0
}

func isValid(s string) bool {
	closingParMatch := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := NewStack[rune]()

	for _, char := range s {
		// check if it's an opening parenthesis or not
		par, isClosingParenthesis := closingParMatch[char]

		if isClosingParenthesis {
			// the top of the stack has to be an opening parenthesis of the same kind
			lastChar := stack.Pop()
			if par != lastChar {
				return false
			}
		} else {
			stack.Push(char)
		}
	}

	return stack.IsEmpty()
}
