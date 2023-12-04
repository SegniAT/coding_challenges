package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	lastIndex := make(map[rune]int)
	left, right := 0, 0
	maxLength := 0.0

	for right < length {
		rightChar := s[right]
		if lastInd, ok := lastIndex[rune(rightChar)]; ok && left <= lastInd {
			left = lastInd + 1
		}

		lastIndex[rune(rightChar)] = right
		maxLength = math.Max(maxLength, float64(right-left+1))
		right++
	}

	return int(maxLength)
}
