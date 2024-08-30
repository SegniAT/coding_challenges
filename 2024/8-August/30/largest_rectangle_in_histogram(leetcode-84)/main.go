package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}

type Value struct {
	ind int
	val int
}

func largestRectangleArea(heights []int) int {
	N := len(heights)
	nextLower, prevLower := make([]int, N), make([]int, N)

	for i := range N {
		nextLower[i] = N
		prevLower[i] = -1
	}

	// nextLower
	// increasing monotonic stack
	stack := []Value{}

	for i, height := range heights {
		for len(stack) > 0 && stack[len(stack)-1].val > height {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextLower[top.ind] = i
		}

		stack = append(stack, Value{ind: i, val: height})
	}

	stack = []Value{}
	fmt.Println("nextLower: ", nextLower)

	// prevLower
	// increasing monotonic stack still but iterate in reverse direction,
	for i := N - 1; i > -1; i-- {
		height := heights[i]
		for len(stack) > 0 && stack[len(stack)-1].val > height {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevLower[top.ind] = i
		}

		stack = append(stack, Value{ind: i, val: height})
	}

	fmt.Println("prevLower: ", prevLower)

	maxRect := 0

	for i := 0; i < N; i++ {
		maxRect = max(maxRect, heights[i]*(nextLower[i]-prevLower[i]-1))
	}

	return maxRect
}
