package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSteps(3))
}

func minSteps(n int) int {

	var helper func(int, int, int) int
	helper = func(lastCopy, count, steps int) int {
		if count > n {
			return math.MaxInt
		}

		if count == n {
			return steps
		}

		// use lastCopy
		opt1 := math.MaxInt
		if lastCopy != 0 {
			opt1 = helper(lastCopy, count+lastCopy, steps+1)
		}

		// new copy
		newCopy := count
		opt2 := helper(newCopy, count+newCopy, steps+2)

		return min(opt1, opt2)
	}

	return helper(0, 1, 0)
}
