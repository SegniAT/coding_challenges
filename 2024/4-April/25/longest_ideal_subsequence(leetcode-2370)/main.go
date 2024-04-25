package main

import "math"

func main() {
}

func longestIdealString(s string, k int) int {
	// index: max
	memo := make(map[int]int)

	for i := range len(s) {
		helper(s, k, i, &memo)
	}

	longest := 0

	for _, val := range memo {
		if val > longest {
			longest = val
		}
	}

	return longest
}

func helper(s string, k int, index int, memo *map[int]int) int {
	if val, ok := (*memo)[index]; ok {
		return val
	}

	res := 1
	longestNext := 0
	for i := index + 1; i < len(s); i++ {
		if difference(rune(s[index]), rune(s[i])) > k {
			continue
		}
		longestNext = int(math.Max(float64(longestNext), float64(helper(s, k, i, memo))))
	}

	(*memo)[index] = res + longestNext

	return res + longestNext
}

func difference(a, b rune) int {
	return int(math.Abs(float64(b - a)))
}
