package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	word1Len, word2Len := len(word1), len(word2)
	result := ""

	word1Turn := true
	w1, w2 := 0, 0

	for w1 < word1Len || w2 < word2Len {
		if word1Turn && w1 < word1Len {
			result += string(word1[w1])
			w1++
			if w2 < word2Len {
				word1Turn = false
			}
		} else if w2 < word2Len {
			result += string(word2[w2])
			w2++
			if w1 < word1Len {
				word1Turn = true
			}
		}

	}

	return result
}
