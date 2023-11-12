package main

import "math"

func main() {

}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	maxCount := 0

	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		if _, ok := set[num-1]; ok {
			continue
		} else {
			count := 1
			currNum := num + 1
			_, ok := set[currNum]
			for ok {
				count++
				currNum++
				_, ok = set[currNum]
			}
			maxCount = int(math.Max(float64(count), float64(maxCount)))
		}
	}

	return maxCount
}
