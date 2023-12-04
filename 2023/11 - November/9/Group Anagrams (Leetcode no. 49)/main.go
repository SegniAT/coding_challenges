package main

import (
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	sortedStrs := make([]string, length)
	groupsIndexMap := make(map[string][]string)

	for ind, str := range strs {
		sortedStrs[ind] = sortString(str)
	}

	for ind, sortedStr := range sortedStrs {
		if indexes, ok := groupsIndexMap[sortedStr]; ok {
			groupsIndexMap[sortedStr] = append(indexes, strs[ind])
		} else {
			groupsIndexMap[sortedStr] = []string{strs[ind]}
		}
	}

	result := make([][]string, len(groupsIndexMap))

	index := 0
	for _, value := range groupsIndexMap {
		result[index] = value
		index++
	}

	return result
}

func sortString(str string) string {
	chars := []rune(str)

	sort.SliceStable(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
