package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	shortestStr := strs[shortestStringIdx(strs)]
	fmt.Println("shortestStr: ", shortestStr)
	lastSimilarCharInd := 0

	for idx, char := range shortestStr {
		for _, str := range strs {
			char2 := str[idx]
			if char != int32(char2) {
				return shortestStr[:lastSimilarCharInd]
			}
		}
		lastSimilarCharInd++
	}
	return shortestStr[:lastSimilarCharInd]
}

func shortestStringIdx(strs []string) int {
	minLen := math.Inf(1)
	minIdx := 0

	for idx, str := range strs {
		strLen := len(str)
		if float64(strLen) < minLen {
			minIdx = idx
			minLen = float64(strLen)
		}
	}
	return minIdx
}
