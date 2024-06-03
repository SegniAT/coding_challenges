package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println()
	fmt.Println(compress([]byte{'a'}))
	fmt.Println()
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	fmt.Println()
	fmt.Println(compress([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}))
	fmt.Println()
	fmt.Println(compress([]byte{'a', 'a', 'a', 'a', 'a', 'a'}))
}

func compress(chars []byte) int {
	chars = append(chars, ' ')
	compressedStr := []string{}

	left, right := 0, 1
	currChar := chars[0]

	for right < len(chars) {
		if chars[right] != currChar {
			compressedStr = append(compressedStr, string(currChar))
			if count := right - left; count > 1 {
				compressedStr = append(compressedStr, fmt.Sprint(count))
			}
			currChar = chars[right]
			left = right
		}

		right++
	}

	cmpStr := strings.Join(compressedStr, "")
	fmt.Println(cmpStr)

	return len(cmpStr)
}
