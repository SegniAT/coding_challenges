package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	res := []string{}

	var helper func(string, int, int)
	helper = func(curr string, openingCount, closingCount int) {
		if openingCount == n && closingCount == n {
			res = append(res, curr)
			return
		}

		if openingCount < n {
			helper(fmt.Sprintf("%s(", curr), openingCount+1, closingCount)
		}

		if openingCount > closingCount {
			helper(fmt.Sprintf("%s)", curr), openingCount, closingCount+1)
		}
	}

	helper("", 0, 0)

	return res
}
