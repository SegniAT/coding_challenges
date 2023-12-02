package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	const MAX_REDS, MAX_GREENS, MAX_BLUES = 12, 13, 14

	possibleGames := 0

	games := []string{}
	for _, game := range lines {
		colon := strings.Index(game, ":")
		game := strings.TrimSpace(game[colon+1:])
		games = append(games, game)
	}

	for ind, game := range games {
		gameNum := ind + 1
		gameSets := strings.Split(game, ";")
		possible := true

		for _, gameSet := range gameSets {
			colors := strings.Split(gameSet, ",")

			for _, color := range colors {
				color = strings.TrimSpace(color)
				countColor := strings.Split(color, " ")

				count, err := strconv.Atoi(countColor[0])
				if err != nil {
					panic(err)
				}
				colorName := countColor[1]

				if colorName == "red" && count > MAX_REDS {
					possible = false
					break
				}
				if colorName == "green" && count > MAX_GREENS {
					possible = false
					break
				}
				if colorName == "blue" && count > MAX_BLUES {
					possible = false
					break
				}
			}

			if !possible {
				break
			}
		}

		if possible {
			possibleGames += gameNum
		}
	}

	fmt.Println("POSSIBLE GAMES: ", possibleGames)

}

func readInput(src string) ([]string, error) {
	res := []string{}

	file, err := os.Open(src)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res, nil
}
