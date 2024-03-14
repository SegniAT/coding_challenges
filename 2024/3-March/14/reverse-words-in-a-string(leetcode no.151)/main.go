package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("   hello world "))
	fmt.Println(reverseWords("a good    example"))
}

func reverseWords(s string) string {
	words := splitString(s)

	l, r := 0, len(words)-1

	for r > l {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}

	return strings.Join(words, " ")
}

func splitString(s string) []string {
	s = strings.Trim(s, " ")
	words := []string{}

	currentWord := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord = fmt.Sprintf("%s%c", currentWord, s[i])
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
