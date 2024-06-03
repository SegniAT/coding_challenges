package main

func main() {
}

func maxDepth(s string) int {
	currD, maxD := 0, 0
	for _, char := range s {
		if char == '(' {
			currD++
			maxD = max(currD, maxD)
		} else if char == ')' {
			currD--
		}
	}

	return maxD
}
