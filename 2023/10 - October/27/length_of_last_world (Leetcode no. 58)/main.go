package main

func main() {

}

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i > -1; i-- {
		currChar := s[i]
		if length == 0 && currChar == ' ' {
			continue
		} else if currChar == ' ' {
			break
		} else {
			length++
		}
	}

	return length
}
