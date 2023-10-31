package main

import "fmt"

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	alphanumeric := ""

	for _, char := range s {
		if isNumber(char) {
			alphanumeric += string(char)
		}

		if isAlphabet(char) {
			alphanumeric += string(toLowerCase(char))
		}
	}

	left, right := 0, len(alphanumeric)-1
	for left < right {
		if alphanumeric[left] != alphanumeric[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func toLowerCase(char rune) rune {
	if !isAlphabet(char) {
		panic("'char' is not an alphabet")
	}

	if 'a' <= char && char <= 'z' {
		return char
	} else {
		diff := char - 'A'
		return 'a' + diff
	}
}

func isAlphabet(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z'
}

func isNumber(char rune) bool {
	return '0' <= char && char <= '9'
}
