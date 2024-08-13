package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	fmt.Println(combinationSum2(candidates, target))
}

// Backtracking problem
func combinationSum2(candidates []int, target int) [][]int {
	N := len(candidates)
	sort.Ints(candidates)

	combinations := [][]int{}

	var helper func(int, int, []int)
	helper = func(currTarget, ind int, sub []int) {
		if currTarget < 0 {
			return
		}

		if currTarget == 0 {
			combo := make([]int, len(sub))
			copy(combo, sub)
			combinations = append(combinations, combo)
			return
		}

		for i := ind; i < N; i++ {
			if i > ind && candidates[i] == candidates[i-1] {
				// Skip duplicates
				continue
			}

			// Take the candidate and move to the next index
			sub = append(sub, candidates[i])
			helper(currTarget-candidates[i], i+1, sub)
			// Backtrack to try other combinations
			sub = sub[:len(sub)-1]
		}
	}

	helper(target, 0, []int{})

	return combinations
}
