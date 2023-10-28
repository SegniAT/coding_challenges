package main

import "fmt"

func main() {
	haystack := "leetcode"
	needle := "leeto"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	left, right := 0, 0
	needleLen := len(needle)

	for right < len(haystack) {
		relativeIdx := right - left
		fmt.Println(haystack[right], needle[relativeIdx])
		if haystack[right] != needle[relativeIdx] {
			right++
			left = right
		} else {
			if relativeIdx == needleLen-1 {
				return left
			}
			right++
		}
	}
	return -1
}
