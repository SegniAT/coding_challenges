package main

func main() {
}

type Value struct {
	ind  int
	char rune
}

func minRemoveToMakeValid(s string) string {
	stack := []Value{}
	skippedIndexes := make(map[int]bool)
	res := ""

	for index, char := range s {
		if char == '(' {
			stack = append(stack, Value{ind: index, char: '('})
		} else if char == ')' {
			if len(stack) == 0 {
				skippedIndexes[index] = true
			} else {
				// pop
				stack = stack[:len(stack)-1]
			}
		}
	}

	for _, val := range stack {
		skippedIndexes[val.ind] = true
	}

	for ind, char := range s {
		if !skippedIndexes[ind] {
			res += string(char)
		}
	}

	return res
}
