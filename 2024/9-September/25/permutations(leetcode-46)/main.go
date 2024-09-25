package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	helper(nums, []int{}, &res, used)
	return res
}

func helper(nums []int, curr []int, res *[][]int, used []bool) {
	if len(nums) == len(curr) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		(*res) = append((*res), tmp)
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}
		used[i] = true
		curr = append(curr, num)
		helper(nums, curr, res, used)
		used[i] = false
		curr = curr[:len(curr)-1]
	}
}
