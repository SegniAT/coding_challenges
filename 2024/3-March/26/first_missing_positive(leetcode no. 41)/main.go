package main

import "math"

func main() {
}

// "brute force"
func firstMissingPositive(nums []int) int {
	hash := make(map[int]bool)

	for _, num := range nums {
		if num > 0 {
			hash[num] = true
		}
	}

	for i := 1; i <= math.MaxInt; i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}

	return -1
}
