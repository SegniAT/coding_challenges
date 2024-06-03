package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	sCount := make(map[byte]int)
	tCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sCount[s[i]]++
		tCount[t[i]]++
	}

	for sChar, sCnt := range sCount {
		if tCnt := tCount[sChar]; sCnt != tCnt {
			return false
		}
	}

	return true
}
