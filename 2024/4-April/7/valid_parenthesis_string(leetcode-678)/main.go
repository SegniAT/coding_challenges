package main

func main() {

}

func checkValidString(s string) bool {
	// memo[index][openingParCount]
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}

	for row := range memo {
		for col := range memo[row] {
			memo[row][col] = -1
		}
	}

	return helper(s, 0, 0, &memo)
}

func helper(s string, index, openingParCount int, memo *[][]int) bool {
	if index == len(s) {
		return openingParCount == 0
	}

	if mem := (*memo)[index][openingParCount]; mem != -1 {
		return mem == 1
	}

	currChar := s[index]
	res := false

	if currChar == '(' {
		res = res || helper(s, index+1, openingParCount+1, memo)
	} else if currChar == ')' {
		if openingParCount > 0 {
			res = res || helper(s, index+1, openingParCount-1, memo)
		}
	} else {
		// * case, as '(', ')' and ''
		res = res || helper(s, index+1, openingParCount+1, memo)
		if openingParCount > 0 {
			res = res || helper(s, index+1, openingParCount-1, memo)
		}
		res = res || helper(s, index+1, openingParCount, memo)
	}

	if res == true {
		(*memo)[index][openingParCount] = 1
	} else {
		(*memo)[index][openingParCount] = 0
	}

	return res
}
