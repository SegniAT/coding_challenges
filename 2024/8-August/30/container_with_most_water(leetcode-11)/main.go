package main

func main() {
}

// Greedy...
func maxArea(height []int) int {
	N := len(height)
	res := -1
	left, right := 0, N-1

	for left < right {
		currArea := min(height[left], height[right]) * (right - left)
		res = max(res, currArea)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return res
}
