package main

import "math"

func main() {

}

func jump(nums []int) int {
	length := len(nums)
	minJumps := make([]float64, length)

	for i := 1; i < length; i++ {
		minJumps[i] = math.Inf(1)
	}

	for i := 0; i < length-1; i++ {
		maxJump := nums[i]
		for jumps := 1; jumps <= maxJump && i+jumps < length; jumps++ {
			minJumps[i+jumps] = math.Min(minJumps[i]+1, minJumps[i+jumps])
		}
	}
	return int(minJumps[length-1])
}
