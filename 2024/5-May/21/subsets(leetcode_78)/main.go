package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{1, 2, 3}
	fmt.Println("input: ", input)
	fmt.Println(subsets1(input))
	fmt.Println(subsets(input))
}

// bitmask method
func subsets1(nums []int) [][]int {
	out := [][]int{}
	ln := len(nums)
	for i := 0; i < int(math.Pow(2, float64(ln))); i++ {
		currNum := i
		ind := 0
		arr := []int{}
		for currNum > 0 {
			if currNum&1 == 1 {
				arr = append(arr, nums[ind])
			}
			currNum = currNum >> 1
			ind++
		}
		out = append(out, arr)
	}

	return out
}

// backtracking
func subsets(nums []int) [][]int {
	out := [][]int{}
	backtrack(nums, &out, []int{}, 0)
	return out
}

func backtrack(nums []int, subsets *[][]int, curr []int, start int) {
	subsetCopy := make([]int, len(curr))
	copy(subsetCopy, curr)
	*subsets = append(*subsets, subsetCopy)

	for i := start; i < len(nums); i++ {
		//take
		curr = append(curr, nums[i])
		backtrack(nums, subsets, curr, i+1)
		// don't take
		curr = curr[:len(curr)-1]
	}
}
