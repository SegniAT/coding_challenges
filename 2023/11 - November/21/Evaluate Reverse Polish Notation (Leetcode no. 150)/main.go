package main

import (
	"fmt"
	"strconv"
)

func main() {

}

type Stack[T any] struct {
	Values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) Push(value T) {
	stack.Values = append(stack.Values, value)
}

func (stack *Stack[T]) Top() (T, bool) {
	var x T
	if len(stack.Values) > 0 {
		x = stack.Values[len(stack.Values)-1]
		return x, true
	}
	return x, false
}

func (stack *Stack[T]) Pop() (T, bool) {
	var x T
	if len(stack.Values) > 0 {
		x, stack.Values = stack.Values[len(stack.Values)-1], stack.Values[:len(stack.Values)-1]
		return x, true
	}
	return x, false
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.Values) == 0
}

func evalRPN(tokens []string) int {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	stack := NewStack[int]()

	for _, token := range tokens {
		_, isOperator := operators[token]

		if isOperator {
			secondOperand, secondOperandExists := stack.Pop()
			firstOperand, firstOperandExists := stack.Pop()

			if secondOperandExists && firstOperandExists {
				switch token {
				case "+":
					stack.Push(firstOperand + secondOperand)
				case "-":
					stack.Push(firstOperand - secondOperand)
				case "*":
					stack.Push(firstOperand * secondOperand)
				case "/":
					stack.Push(firstOperand / secondOperand)
				}
			} else {
				return -1
			}

		} else {
			tokenInt, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println(fmt.Errorf("strconv: %d", err))
				return -1
			}
			stack.Push(tokenInt)
		}
	}

	result, resultExists := stack.Pop()

	if !resultExists {
		return -1
	}

	return result
}
