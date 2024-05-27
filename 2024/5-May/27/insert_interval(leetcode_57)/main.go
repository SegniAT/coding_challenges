package main

import (
	"fmt"
	"slices"
)

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}

	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	slices.SortFunc(intervals, func(a []int, b []int) int {
		if a[0] < b[0] {
			return -1
		} else {
			return 1
		}
	})

	out := [][]int{}

	out = append(out, intervals[0])

	for i := 1; i < len(intervals); i++ {
		outLn := len(out)
		interval := intervals[i]
		latestOut := out[outLn-1]

		if latestOut[1] >= interval[0] {
			if latestOut[1] < interval[1] {
				latestOut[1] = interval[1]
			}
		} else {
			out = append(out, interval)
		}

	}

	return out
}
