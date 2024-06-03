package main

import (
	"fmt"
	"math"
)

func main() {
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = math.MaxInt
	}

	queue := []int{}
	queue = append(queue, n)

	steps := 0
	for len(queue) > 0 {
		// process level
		for levelSize := len(queue); levelSize > 0; levelSize-- {
			// pop
			popped := queue[0]
			queue = queue[1:]

			if dp[popped] <= steps {
				continue
			}

			dp[popped] = steps

			// add children
			for i := int(math.Sqrt(float64(popped))); i > 0; i-- {
				child := popped - (i * i)
				queue = append(queue, child)
			}

		}
		steps++
	}

	return dp[0]
}
