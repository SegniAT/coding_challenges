package main

import (
	"fmt"
	"slices"
)

func main() {
}

func findRelativeRanks(score []int) []string {
	sortedScore := []int{}
	for _, sc := range score {
		sortedScore = append(sortedScore, sc)
	}

	slices.Sort(sortedScore)

	rank := make(map[int]int)

	ln := len(sortedScore)
	for ind, sc := range sortedScore {
		rank[sc] = ln - ind
	}

	ans := []string{}

	for _, sc := range score {
		scoreRank := rank[sc]
		ans = append(ans, rankFromIndex(scoreRank))
	}

	return ans
}

func rankFromIndex(ind int) string {
	switch ind {
	case 1:
		return "Gold Medal"
	case 2:
		return "Silver Medal"
	case 3:
		return "Bronze Medal"
	default:
		return fmt.Sprintf("%d", ind)
	}
}
