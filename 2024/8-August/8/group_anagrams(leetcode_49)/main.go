package main

import (
	"sort"
	"strings"
)

func main() {
}

func groupAnagrams(strs []string) [][]string {
	cache := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		cache[sortedStr] = append(cache[sortedStr], str)
	}

	res := [][]string{}

	for _, val := range cache {
		res = append(res, val)
	}

	return res
}

func sortString(str string) string {
	arr := strings.Split(str, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}
