package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strLen := len(s)
	res := ""
	offset1, offset2 := numRows*2-2, 0
	for row := 0; row < numRows; row++ {
		currInd := row

		if currInd < strLen {
			res += string(s[currInd])
		}

		for currInd < strLen && currInd > -1 {
			currInd += offset1
			if currInd >= strLen {
				continue
			}

			if offset1 > 0 {
				res += string(s[currInd])
			}

			currInd += offset2
			if currInd >= strLen {
				continue
			}

			if offset2 > 0 {
				res += string(s[currInd])
			}
		}

		offset1 -= 2
		offset2 += 2
	}

	return res
}
