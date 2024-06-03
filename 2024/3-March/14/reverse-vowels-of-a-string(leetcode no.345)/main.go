package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	sElements := strings.Split(s, "")
	left, right := 0, len(s)-1

	for right > left {
		leftIsVowel, rightIsVowel := isVowel(sElements[left]), isVowel(sElements[right])
		if leftIsVowel && rightIsVowel {
			sElements[left], sElements[right] = sElements[right], sElements[left]
			left++
			right--
		} else if leftIsVowel {
			right--
		} else if rightIsVowel {
			left++
		} else {
			left++
			right--
		}
	}

	return strings.Join(sElements, "")
}

func isVowel(char string) bool {
	uppercase := []string{"A", "E", "I", "O", "U"}
	lowecase := []string{"a", "e", "i", "o", "u"}

	for _, c := range uppercase {
		if char == c {
			return true
		}
	}

	for _, c := range lowecase {
		if char == c {
			return true
		}
	}

	return false
}
