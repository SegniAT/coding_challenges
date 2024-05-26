package main

import (
	"fmt"
	"slices"
)

func main() {
	inp := [][]int{{2, 6}, {1, 3}, {15, 18}, {8, 10}}
	// sorted: [[1 3] [2 6] [8 10] [15 18]]

	fmt.Println(merge(inp))

	inp = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(inp))

	fmt.Println(merge([][]int{{1, 2}}))
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(i, j []int) int {
		if i[0] > j[0] {
			return 1
		} else if i[0] < j[0] {
			return -1
		} else {
			return 0
		}
	})

	fmt.Println("sorted: ", intervals)

	var out [][]int

	out = append(out, intervals[0])

	for i := 1; i < len(intervals); i++ {
		outLen := len(out)
		interval := intervals[i]

		if out[outLen-1][1] >= interval[0] {
			if out[outLen-1][1] < interval[1] {
				out[outLen-1][1] = interval[1]
			}
		} else {
			out = append(out, intervals[i])
		}
	}

	return out
}
