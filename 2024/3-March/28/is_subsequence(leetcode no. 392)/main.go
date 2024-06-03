package main

func main() {
}

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	if sLen > tLen {
		return false
	}

	sPtr, tPtr := 0, 0

	for ; tPtr < tLen; tPtr++ {
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
	}

	return sPtr == len(s)
}
