package main

import (
	"fmt"
	"slices"
)

func main() {
	c := []int{10, 1, 2, 7, 6, 1, 5}
	t := 8
	fmt.Println(combinationSum2(c, t))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	slices.Sort(candidates)
	helper(candidates, target, []int{}, &res, 0)
	return res
}

func helper(c []int, t int, curr []int, res *[][]int, ind int) {
	if t < 0 {
		return
	}

	if t == 0 {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
		return
	}

	for i := ind; i < len(c); i++ {
		if i > ind && c[i-1] == c[i] {
			continue
		}
		curr = append(curr, c[i])
		helper(c, t-c[i], curr, res, i+1)
		curr = curr[:len(curr)-1]
	}
}
