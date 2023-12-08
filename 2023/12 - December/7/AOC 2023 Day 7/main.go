package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	card string
	bid  int
}

type CardType int

const (
	HighCard     CardType = iota // 0
	OnePair                      // 1
	TwoPair                      // 2
	ThreeOfAKind                 // 3
	FullHouse                    // 4
	FourOfAKind                  // 5
	FiveOfAKind                  // 6
)

var cards = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
var cardStrength = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

type ByStrength []*Hand

func (bs ByStrength) Len() int      { return len(bs) }
func (bs ByStrength) Swap(i, j int) { bs[i], bs[j] = bs[j], bs[i] }
func (bs ByStrength) Less(i, j int) bool {
	handOne, handTwo := bs[i], bs[j]

	handOneTypeStronger, handTwoTypeStronger := makeHandStrongerBetter(handOne.card), makeHandStrongerBetter(handTwo.card)

	handOneType, handTwoType := getCardType(handOneTypeStronger), getCardType(handTwoTypeStronger)

	if handOneType != handTwoType {
		return handOneType < handTwoType
	} else {
		for i := 0; i < 5; i++ {
			cardOne, cardTwo := rune(handOne.card[i]), rune(handTwo.card[i])
			if cardStrength[cardOne] > cardStrength[cardTwo] {
				return false
			} else if cardStrength[cardOne] < cardStrength[cardTwo] {
				return true
			}
		}
		return true
	}
}

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	for _, in := range input {
		fmt.Println(in)
	}

	fmt.Println()
	fmt.Println()

	sort.Sort(ByStrength(input))

	totalWinnings := 0

	fmt.Println("SORTED:")
	for ind, in := range input {
		totalWinnings += (ind + 1) * in.bid
		fmt.Println(in)
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Total winnings: ", totalWinnings)
	// sample output: 5905
	// first attempt: 249959407 (too high) ("better" function)
	// second attempt: 249247101 (too low) (first function)
	// third attempt: 249227780 (too low) (first function)
	// fourth attempt: 249817836 ("better" function) boooyah!!!!!!!!!

	// for _, inp := range input {
	// 	fmt.Println(inp.card, makeHandStronger(inp.card))
	// }
	// fmt.Println()
	// for _, inp := range input {
	// 	fmt.Println(inp.card, makeHandStrongerBetter(inp.card))
	// }

}

func getCardType(hand string) CardType {
	count := make(map[rune]int)

	for _, card := range hand {
		if _, ok := count[card]; !ok {
			count[card] = 1
		} else {
			count[card]++
		}
	}

	distinctCards := len(count)

	switch distinctCards {
	case 1:
		return FiveOfAKind
	case 2:
		for _, cnt := range count {
			if cnt == 4 {
				return FourOfAKind
			} else if cnt == 3 {
				return FullHouse
			}
		}
	case 3:
		for _, cnt := range count {
			if cnt == 3 {
				return ThreeOfAKind
			}
		}

		return TwoPair
	case 4:
		for _, cnt := range count {
			if cnt == 2 {
				return OnePair
			}
		}
	}

	return HighCard
}

// DOESN'T WORK!!!
func makeHandStronger(hand string) string {
	if strings.Index(hand, "J") == -1 {
		return hand
	}

	count := make(map[rune]int)

	for _, card := range hand {
		if _, ok := count[card]; !ok {
			count[card] = 1
		} else {
			count[card]++
		}
	}

	oldHandType := getCardType(hand)

	switch oldHandType {
	case FiveOfAKind:
		return hand
	case FourOfAKind:
		// four same, one different. Try to make the one like others if it's J
		otherCard := '.'
		for card, cnt := range count {
			if cnt == 4 {
				otherCard = card
				break
			}
		}

		if count['J'] == 1 {
			hand = strings.Replace(hand, "J", string(otherCard), 1)
		}
	case FullHouse:
		// three cards same, remaining two are same
		otherCard := '.'
		for card := range count {
			if card != 'J' {
				otherCard = card
				break
			}
		}

		hand = strings.ReplaceAll(hand, "J", string(otherCard))
	case ThreeOfAKind:
		// three cards the same, the remaining two are different
		// 2 possiblities
		if count['J'] == 3 {
			otherCard := '.'
			for card := range count {
				if card != 'J' {
					otherCard = card
					break
				}
			}
			hand = strings.ReplaceAll(hand, "J", string(otherCard))
		} else {
			// count['J'] == 1
			// change 'J' to the card with count 3, changes it to 'Four of a kind'
			otherCard := '.'
			for card, cnt := range count {
				if cnt == 3 {
					otherCard = card
					break
				}
			}

			hand = strings.ReplaceAll(hand, "J", string(otherCard))
		}
	case TwoPair:
		// two cards are the same, two other cards are same, one unique
		if count['J'] == 2 {
			otherCard := '.'
			for card, cnt := range count {
				if card != 'J' && cnt == 2 {
					otherCard = card
					break
				}
			}

			hand = strings.ReplaceAll(hand, "J", string(otherCard))
		} else {
			// count['J'] == 1
			otherCard := '.'
			for card, cnt := range count {
				if cnt == 2 {
					otherCard = card
					break
				}
			}

			hand = strings.Replace(hand, "J", string(otherCard), 1)
		}
	case OnePair:
		// two cards are the same, others are all unique
		if count['J'] == 2 {
			otherCard := '.'
			for card := range count {
				if card != 'J' {
					otherCard = card
					break
				}
			}

			hand = strings.ReplaceAll(hand, "J", string(otherCard))
		}
	case HighCard:
		// all unique
		otherCard := '.'
		for card := range count {
			if card != 'J' {
				otherCard = card
				break
			}
		}

		hand = strings.Replace(hand, "J", string(otherCard), 1)
	}

	return hand
}

func makeHandStrongerBetter(hand string) string {
	if strings.Index(hand, "J") == -1 {
		return hand
	}

	count := make(map[rune]int)

	for _, card := range hand {
		if _, ok := count[card]; !ok {
			count[card] = 1
		} else {
			count[card]++
		}
	}

	mostCard, mostCardCount := '.', 0

	for card, cnt := range count {
		if card != 'J' && cnt > mostCardCount {
			mostCard = card
			mostCardCount++
		}
	}

	if mostCardCount == 0 {
		return hand
	}

	hand = strings.ReplaceAll(hand, "J", string(mostCard))
	return hand
}

func loadInput(src string) ([]*Hand, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	out := []*Hand{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		card := line[0]
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		out = append(out, &Hand{
			card: card,
			bid:  bid,
		})
	}

	return out, nil
}
