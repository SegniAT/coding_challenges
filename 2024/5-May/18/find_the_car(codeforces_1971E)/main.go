package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	n, k, q   int
	distances []int
	times     []int
	questions []int
}

func main() {
	tests, err := loadInput()
	if err != nil {
		log.Fatalln("unable to read file: ", err)
	}

	//fmt.Println(tests)

	outputs := [][]int{}

	for _, test := range tests {
		out := []int{}
		for _, distance := range test.questions {
			out = append(out, solve(test.distances, test.times, distance))
		}
		outputs = append(outputs, out)
	}

	print2Dgrid(outputs)

	/*
		testInd := 0
		fmt.Println(tests[testInd])
		//fmt.Println(solve(tests[testInd].distances, tests[testInd].times, tests[testInd].questions[0]))
		fmt.Println(solve(tests[testInd].distances, tests[testInd].times, tests[testInd].questions[1]))
		//fmt.Println(solve(tests[testInd].distances, tests[testInd].times, tests[testInd].questions[2]))
		//fmt.Println(solve(tests[testInd].distances, tests[testInd].times, tests[testInd].questions[3]))

	*/
}

func print2Dgrid(arr [][]int) {
	for _, row := range arr {
		strArr := ""
		for _, col := range row {
			strArr += fmt.Sprintf("%d", col)
			strArr += " "
		}

		fmt.Println(strArr)
	}
}

func solve(distances []int, times []int, distance int) int {
	//fmt.Println("distance: ", distance)
	left, right := 0, len(distances)-1

	for left < right {
		mid := (right + left) / 2
		if distances[mid] >= distance {
			right = mid
			if distances[mid] == distance {
				break
			}
		} else {
			left = mid + 1
		}
	}

	//fmt.Println("right:", right)

	if distance == distances[right] {
		return times[right]
	}

	if right == 0 {
		speed := float64(distances[0]) / float64(times[0])
		return int(float64(distance) / speed)
	}

	speed := float64(distances[right]-distances[right-1]) / float64(times[right]-times[right-1])
	timeFromPrevPoint := float64(distance-distances[right-1]) / speed

	return times[right-1] + int(timeFromPrevPoint)
}

func loadInput() ([]Input, error) {
	out := []Input{}
	scanner := bufio.NewScanner(os.Stdin)

	// test cases
	scanner.Scan()
	tests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	for tests > 0 {
		// n, k, q
		scanner.Scan()
		line1 := strToIntArr(scanner.Text())
		inp := Input{
			n: line1[0],
			k: line1[1],
			q: line1[2],
		}

		// ai
		scanner.Scan()
		inp.distances = strToIntArr(scanner.Text())

		// bi
		scanner.Scan()
		inp.times = strToIntArr(scanner.Text())

		for range inp.q {
			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			inp.questions = append(inp.questions, num)
		}

		out = append(out, inp)
		tests--
	}

	return out, nil
}

func strToIntArr(inp string) []int {
	arr := strings.Split(inp, " ")
	out := []int{}
	for _, str := range arr {
		num, _ := strconv.Atoi(str)
		out = append(out, num)
	}
	return out
}
