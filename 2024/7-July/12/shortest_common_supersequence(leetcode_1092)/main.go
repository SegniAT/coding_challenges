package main

import "fmt"

// WRONG!!! skipping

func main() {
	str1, str2 := "abac", "cab"
	fmt.Println(shortestCommonSupersequence(str1, str2))
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	len1, len2 := len(str1), len(str2)
	dimension := len1 + len2
	dp := make([][]string, dimension)
	for row := range dimension {
		dp[row] = make([]string, dimension)
		for col := range dimension {
			dp[row][col] = ""
		}
	}

	res := helper(&dp, str1, str2, 0, 0)
	fmt.Println("dp: ", dp)
	return res
}

func helper(dp *[][]string, str1, str2 string, ind1, ind2 int) string {
	len1, len2 := len(str1), len(str2)

	if ind1 >= len1 && ind2 >= len2 {
		return maximum(str1, str2)
	}

	if ind1 >= len1 {
		fmt.Println(ind1, ind2, str1, str2)
		return fmt.Sprintf("%s%s", str1, string(str2[ind2]))
	}

	if ind2 >= len2 {
		fmt.Println(ind1, ind2, str1, str2)
		return fmt.Sprintf("%s%s", str2, string(str1[ind1]))
	}

	if res := (*dp)[ind1][ind2]; res != "" {
		return res
	}

	var res string
	if str1[ind1] == str2[ind2] {
		res = helper(dp, str1, str2, ind1+1, ind2+1)
	} else {
		// insert to first string
		inserted := insertAt(str1, string(str2[ind2]), ind1)
		res = helper(dp, inserted, str2, ind1+1, ind2+1)

		// insert to second string
		inserted = insertAt(str2, string(str1[ind1]), ind2)
		res = minimum(res, helper(dp, str1, inserted, ind1+1, ind2+1))

		// skip first string character
		res = minimum(res, helper(dp, str1, str2, ind1, ind2+1))

		// skip second string character
		res = minimum(res, helper(dp, str1, str2, ind1+1, ind2))

	}

	fmt.Println()
	(*dp)[ind1][ind2] = res
	return res
}

func insertAt(str, val string, ind int) string {
	return fmt.Sprintf("%s%s%s", str[:ind], val, str[ind:])
}

func maximum(str1, str2 string) string {
	if len(str1) > len(str2) {
		return str1
	}

	return str2
}

func minimum(str1, str2 string) string {
	if len(str1) < len(str2) {
		return str1
	}

	return str2
}
