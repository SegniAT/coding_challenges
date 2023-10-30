package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 4}
	fmt.Println(canJump(nums))
}

type STATE int

const (
	NO STATE = iota
	YES
	UNCHECKED
)

func canJump(nums []int) bool {
	state := make([]STATE, len(nums))
	for ind, _ := range state {
		state[ind] = UNCHECKED
	}

	return canJumpHelper(0, nums, state)
}

func canJumpHelper(currInd int, nums []int, state []STATE) bool {
	if currIndState := state[currInd]; currIndState != UNCHECKED {
		return currIndState == YES
	}
	if currInd >= len(nums)-1 {
		return true
	}

	maxJumpDistance := nums[currInd]
	canJumpVar := false
	for i := 1; i <= maxJumpDistance; i++ {
		canJumpVar = canJumpVar || canJumpHelper(currInd+i, nums, state)
	}

	if canJumpVar {
		state[currInd] = YES
	} else {
		state[currInd] = NO
	}

	return canJumpVar
}

func canJumpEasier(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length)
	dp[0] = true

	for i := 0; i < length; i++ {
		if !dp[i] {
			continue
		}

		maxJump := i + nums[i]
		for j := i + 1; j <= maxJump && j < length; j++ {
			dp[j] = true
		}
	}

	return dp[length]
}
