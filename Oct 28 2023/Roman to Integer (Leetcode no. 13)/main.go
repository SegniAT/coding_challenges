package main

func main() {
	//romanToInt("LVIII")
	//romanToInt("III")
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total, strLen := 0, len(s)

	for i := strLen - 1; i > -1; i-- {
		nextCharInd := i + 1
		if nextCharInd < strLen && (romanToIntMap[rune(s[i])] < romanToIntMap[rune(s[nextCharInd])]) {
			total -= romanToIntMap[rune(s[i])]
		} else {
			total += romanToIntMap[rune(s[i])]
		}
	}

	return total
}
