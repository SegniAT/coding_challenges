package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	helper(nums, []int{}, &res, 0)
	return res
}

func helper(nums []int, curr []int, res *[][]int, ind int) {
	temp := make([]int, len(curr))
	copy(temp, curr)
	*res = append(*res, temp)

	for i := ind; i < len(nums); i++ {
		if i > ind && nums[i] == nums[i-1] {
			continue
		}
		curr = append(curr, nums[i])
		helper(nums, curr, res, i+1)
		curr = curr[:len(curr)-1]
	}
}
