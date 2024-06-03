package main

import (
	"fmt"
)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println()

	// no overlaps
	//first
	fmt.Println(insert([][]int{{3, 5}, {6, 7}}, []int{1, 2}))
	//middle
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {13, 16}}, []int{11, 12}))
	//last
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 14}}, []int{99, 999}))

}

func insert(intervals [][]int, newInterval []int) [][]int {
	intLen := len(intervals)
	result := [][]int{}
	inserted := false

	for i := 0; i < intLen; i++ {

		// find x first
		if newInterval[0] <= intervals[i][1] && !inserted {
			var x int
			// out
			if intervals[i][0] > newInterval[0] {
				x = newInterval[0]
			} else { // in
				x = intervals[i][0]
			}

			// find y
			for i < intLen && intervals[i][1] < newInterval[1] {
				i++
			}

			if i >= intLen {
				result = append(result, []int{x, newInterval[1]})
			} else {
				// out
				if intervals[i][0] > newInterval[1] {
					result = append(result, []int{x, newInterval[1]})
					result = append(result, intervals[i])
				} else { // in
					result = append(result, []int{x, intervals[i][1]})
				}
			}
			inserted = true

		} else {
			result = append(result, intervals[i])
		}
	}

	if !inserted {
		result = append(result, newInterval)
	}

	return result
}
