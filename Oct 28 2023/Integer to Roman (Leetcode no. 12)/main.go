package main

import (
	"fmt"
	"strings"
)

func main() {
	//for i := 1; i <= 50; i++ {
	//	fmt.Println("i: ", i, intToRoman(i))
	//}

	fmt.Println("1994", intToRoman(1994))

}

func intToRoman(num int) string {
	regulars := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1_000}
	intRomanMap := map[int]string{
		1:     "I",
		4:     "IV",
		5:     "V",
		9:     "IX",
		10:    "X",
		40:    "XL",
		50:    "L",
		90:    "XC",
		100:   "C",
		400:   "CD",
		500:   "D",
		900:   "CM",
		1_000: "M",
	}

	if roman, ok := intRomanMap[num]; ok {
		return roman
	} else {
		for i := len(regulars) - 1; i > -1; i-- {
			regular := regulars[i]

			if quotient := num / regular; quotient != 0 {
				remainder := num % regular
				if roman, ok := intRomanMap[regular*quotient]; ok {
					return fmt.Sprintf("%s%s", roman, intToRoman(remainder))
				}
				return fmt.Sprintf("%s%s", strings.Repeat(intRomanMap[regular], quotient), intToRoman(remainder))
			}
		}
	}
	return ""
}
