package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))

	fmt.Println(findMinArrowShots([][]int{{1, 2}, {4, 5}, {1, 5}}))
}

func findMinArrowShots(points [][]int) int {
	points = sortPoints(points)
	final := [][]int{}

	// 1 <= points.length <= 10^5 but meh
	if len(points) == 0 {
		return 0
	}

	for _, point := range points {

		if len(final) == 0 || point[0] > final[len(final)-1][1] {
			final = append(final, point)
		} else {
			final[len(final)-1][0] = point[0]
			final[len(final)-1][1] = int(math.Min(float64(final[len(final)-1][1]), float64(point[1])))
		}
	}

	fmt.Println(final)
	return len(final)
}

func sortPoints(points [][]int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][1] < points[j][1]
	})
	return points
}
