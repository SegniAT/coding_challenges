package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	matchingMapST := make(map[rune]rune)
	matchingMapTS := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		sChar, tChar := rune(s[i]), rune(t[i])

		if char, ok := matchingMapST[sChar]; ok && char != tChar {
			return false
		}
		matchingMapST[sChar] = tChar
	}

	for i := 0; i < len(s); i++ {
		sChar, tChar := rune(s[i]), rune(t[i])

		if char, ok := matchingMapTS[tChar]; ok && char != sChar {
			return false
		}
		matchingMapTS[tChar] = sChar
	}

	return true
}
