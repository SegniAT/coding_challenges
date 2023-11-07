package main

import "math"

func main() {

}

func maxAreaNaive(height []int) int {
	max := math.Inf(-1)
	length := len(height)

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			minHeight := math.Min(float64(height[i]), float64(height[j]))
			max = math.Max(max, minHeight*float64(j-i))
		}
	}

	return int(max)
}

// greedy, 2 pointer
func maxArea(height []int) int {
	max, length := math.Inf(-1), len(height)
	left, right := 0, length-1

	for right > left {
		minHeight := math.Min(float64(height[left]), float64(height[right]))
		max = math.Max(max, minHeight*float64(right-left))

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return int(max)
}
