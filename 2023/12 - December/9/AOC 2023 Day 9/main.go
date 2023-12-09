package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	histories := make([][]int, len(input))

	for ind, history := range input {
		histories[ind], err = parseNums(history)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("input: ", histories)
	fmt.Println()

	sum := 0
	for _, history := range histories {
		sequence := generateSequence(history)
		prediction := predictPrev(sequence)
		sum += prediction
	}

	fmt.Println("SUM: ", sum)
}

func generateSequence(history []int) [][]int {
	seq := [][]int{}
	seq = append(seq, history)

	last := history

	for last[len(last)-1] != 0 {
		subSeq := []int{}

		for i := 1; i < len(last); i++ {
			subSeq = append(subSeq, last[i]-last[i-1])
		}

		seq = append(seq, subSeq)

		last = subSeq
	}

	return seq
}

func predictNext(sequence [][]int) int {
	curr := 0

	for i := len(sequence) - 1; i > -1; i-- {
		subSeq := sequence[i]
		curr += subSeq[len(subSeq)-1]
	}

	return curr
}

func predictPrev(sequence [][]int) int {
	curr := 0

	for i := len(sequence) - 1; i > -1; i-- {
		curr = sequence[i][0] - curr
	}

	return curr
}

func parseNums(line string) ([]int, error) {
	numStrs := strings.Split(line, " ")
	out := []int{}

	for _, numStr := range numStrs {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}

		out = append(out, numInt)
	}

	return out, nil
}

func loadInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := []string{}

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, err
}
