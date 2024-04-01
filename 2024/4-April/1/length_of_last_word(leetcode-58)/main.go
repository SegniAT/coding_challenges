package main

func main() {
}

func lengthOfLastWord(s string) int {
	str := " " + s

	word := 0
	for i := len(str) - 1; i > -1; i-- {
		if str[i] != ' ' {
			word++
		} else if word != 0 {
			break
		}
	}

	return word
}
