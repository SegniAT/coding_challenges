package main

import "fmt"

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	left, right, result := 0, 0, ""

	for right < len(s) {
		if s[right] == ' ' {
			if left < right {
				if result == "" {
					result = s[left:right]
				} else {
					result = s[left:right] + " " + result
				}
			}
			right++
			left = right
		} else {
			right++
		}
	}

	if left < right {
		if result == "" {
			result = s[left:right]
		} else {
			result = s[left:right] + " " + result
		}
	}

	return result
}
