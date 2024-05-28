package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {4, 5}, {1, 5}}))
}

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a []int, b []int) int {
		if a[0] <= b[0] {
			return -1
		} else {
			return 1
		}
	})

	//fmt.Println("sorted: ", points)

	out := [][]int{}
	out = append(out, points[0])

	for i := 1; i < len(points); i++ {
		outLatest := out[len(out)-1]
		point := points[i]

		if outLatest[1] < point[0] {
			out = append(out, point)
		} else {
			outLatest[0] = point[0]
			if outLatest[1] > point[1] {
				outLatest[1] = point[1]
			}
		}
	}

	return len(out)
}
