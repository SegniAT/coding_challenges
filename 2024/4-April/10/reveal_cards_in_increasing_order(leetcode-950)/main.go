package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Println(deckRevealedIncreasing([]int{1, 1000}))

}

func deckRevealedIncreasing(deck []int) []int {
	sortedDeck := []int{}
	for _, card := range deck {
		sortedDeck = append(sortedDeck, card)
	}

	slices.Sort(sortedDeck)

	finalIndexPositions := []int{}
	queue := []int{}
	for index := range deck {
		queue = append(queue, index)
	}

	// simulation
	for len(queue) > 0 {
		finalIndexPositions = append(finalIndexPositions, queue[0])
		queue = queue[1:]

		if len(queue) > 0 {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}

	ans := []int{}
	for range len(deck) {
		ans = append(ans, -1)
	}

	for i := range len(deck) {
		insertedNum := sortedDeck[i]
		insertedPosition := finalIndexPositions[i]
		ans[insertedPosition] = insertedNum
	}

	return ans
}
