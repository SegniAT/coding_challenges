package main

import "sort"

func main() {}

func findMaxK(nums []int) int {
	negatives := make(map[int]bool)
	positives := []int{}

	for _, num := range nums {
		if num < 0 {
			negatives[num] = true
		} else {
			positives = append(positives, num)
		}
	}

	sort.Ints(positives)

	for i := len(positives) - 1; i > -1; i-- {
		pos := positives[i]
		if _, ok := negatives[-pos]; ok {
			return pos
		}
	}

	return -1
}
