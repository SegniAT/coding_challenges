package main

import "strings"

func main() {
}

func isPalindrome(s string) bool {
	N := len(s)
	left, right := 0, N-1

	s = strings.ToLower(s)

	for right > left {
		for left < N && !isAlphanumeric(rune(s[left])) {
			left++
		}

		for right > -1 && !isAlphanumeric(rune(s[right])) {
			right--
		}
		if left >= N || right < 0 {

			continue
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || '0' <= char && char <= '9'
}
