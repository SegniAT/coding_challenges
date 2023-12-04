package main

import (
	"bufio"
	"fmt"
	"math"
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

	power := 0

	games := []string{}
	for _, game := range lines {
		colon := strings.Index(game, ":")
		game := strings.TrimSpace(game[colon+1:])
		games = append(games, game)
	}

	for _, game := range games {
		minRed, minGreen, minBlue := 0, 0, 0
		gameSets := strings.Split(game, ";")

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

				if colorName == "red" {
					minRed = int(math.Max(float64(minRed), float64(count)))
				}
				if colorName == "green" {
					minGreen = int(math.Max(float64(minGreen), float64(count)))
				}
				if colorName == "blue" {
					minBlue = int(math.Max(float64(minBlue), float64(count)))
				}
			}

		}

		power += minRed * minGreen * minBlue

	}

	fmt.Println("POWER OF SETS: ", power)

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
