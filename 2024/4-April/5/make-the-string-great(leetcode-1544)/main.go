package main

func main() {
}

func makeGood(s string) string {
	stack := []rune{}

	for i := 0; i < len(s); i++ {
		currChar := rune(s[i])
		if len(stack) == 0 {
			stack = append(stack, currChar)
			continue
		}

		seekTop := stack[len(stack)-1]
		sameLetter, sameCase := sameLetterWithCase(currChar, seekTop)

		if sameLetter && !sameCase {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, currChar)
		}
	}

	res := ""
	for _, char := range stack {
		res += string(char)
	}

	return res
}

func sameLetterWithCase(a, b rune) (bool, bool) {
	if a == b {
		return true, true
	} else if ('a' <= a && a <= 'z') && ('A' <= b && b <= 'Z') { // a lower, b upper
		return (a - 'a') == (b - 'A'), false
	} else if ('A' <= a && a <= 'Z') && ('a' <= b && b <= 'z') { // a upper, b lower
		return (a - 'A') == (b - 'a'), false
	}
	return false, false
}
