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

var cards = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardStrength = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

type ByStrength []*Hand

func (bs ByStrength) Len() int      { return len(bs) }
func (bs ByStrength) Swap(i, j int) { bs[i], bs[j] = bs[j], bs[i] }
func (bs ByStrength) Less(i, j int) bool {
	handOne, handTwo := bs[i], bs[j]

	handOneType, handTwoType := getCardType(handOne.card), getCardType(handTwo.card)

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

	// for _, in := range input {
	// 	fmt.Println(in)
	// }

	// fmt.Println()
	// fmt.Println()

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
