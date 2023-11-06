package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCharCount := make(map[rune]int)
	tCharCount := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		sCharCount[rune(s[i])]++
		tCharCount[rune(t[i])]++
	}

	for _, sChar := range s {
		if sCharCount[sChar] != tCharCount[sChar] {
			return false
		}
	}

	return true
}
