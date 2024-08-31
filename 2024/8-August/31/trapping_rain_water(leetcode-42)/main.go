package main

import "fmt"

func main() {
	height := []int{3, 1, 0, 2}
	fmt.Println(trap(height))
}

// 2 pointer method
/*
func trap(height []int) int {
	N := len(height)
	left, right := 0, N-1
	maxLeft, maxRight := 0, 0

	trappedWater := 0

	// choose the bigger height as a pivot and tighten the left/right

	for left < right {
		maxLeft, maxRight = max(maxLeft, height[left]), max(maxRight, height[right])

		if height[left] < height[right] {
			// if there's a height bigger than this, the current block traps water
			// if not, 0 is added
			trappedWater += maxLeft - height[left]
			left++
		} else {
			trappedWater += maxRight - height[right]
			right--
		}

	}

	return trappedWater
} */

// decreasing monotonic stack
type Height struct {
	ind int
	val int
}

func trap(height []int) int {
	stack := []Height{}
	trappedRain := 0

	for i, h := range height {
		// keep a decreasing order
		for len(stack) > 0 && stack[len(stack)-1].val < h {
			// pop top element
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				newTop := stack[len(stack)-1]

				/*
					 distance is used in the following scenario for example
					 eg. [3,1,0,2]

					| []
					| [] 2 2 []
					|_[][]_1_[]__

					> when trying to add 2 the stack is [3,1,0]
					> on the first pass the water labeled as 1 is added

					> but on the second after doing 2-1=1...we need to multiply that by the distance
					> this is because we did not add water above the first one added

				*/
				distance := i - newTop.ind - 1
				currBlockWater := min(newTop.val, h) - popped.val

				trappedRain += currBlockWater * distance
			}

		}

		stack = append(stack, Height{ind: i, val: h})
	}

	return trappedRain
}
