package main

import "strings"

func main() {

}

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)

	for ind, char := range pattern {
		word := words[ind]

		mappedWord, ok := charToWord[char]

		if ok && word != mappedWord {
			return false
		}

		charToWord[char] = word
	}

	for ind, word := range words {
		char := rune(pattern[ind])

		mappedchar, ok := wordToChar[word]

		if ok && char != mappedchar {
			return false
		}

		wordToChar[word] = char
	}

	return true
}
