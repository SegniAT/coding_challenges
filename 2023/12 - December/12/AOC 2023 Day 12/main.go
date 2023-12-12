package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UNKNOWN     = '?'
	OPERATIONAL = '.'
	DAMAGED     = '#'
)

type Record struct {
	record  string
	damaged []int
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, inp := range input {
		// fmt.Println(inp.record, inp.damaged)
		arrangements := countArrangements(inp.record, inp.damaged)
		sum += arrangements
		// fmt.Println(arrangements)
		// fmt.Println()
	}
	fmt.Println("SUM: ", sum)

	// PART I
	// Attempt 1: 8073 (too high)
	// Attempt 2: 7633

	// PART II
	// Attempt 1: 
}

/*
.??..##.??#??##? [2 2 2 2]
2

?.?#?..##?.?.. [1 2]
3
*/

func countArrangements(record string, damaged []int) int {
	if record == "" {
		if len(damaged) != 0 {
			return 0
		}
		return 1
	}

	if len(damaged) == 0 {
		damagedFound := false
		for i := 0; i < len(record); i++ {
			if record[i] == DAMAGED {
				damagedFound = true
				break
			}
		}

		if damagedFound {
			return 0
		} else {
			return 1
		}
	}

	if record[0] == OPERATIONAL {
		ind := 0
		for ind < len(record) && record[ind] == OPERATIONAL {
			ind++
		}
		return countArrangements(record[ind:], damaged)
	} else if record[0] == DAMAGED {
		if len(record) < damaged[0] {
			return 0
		}
		operationalExists := false
		i := 0
		for ; i < damaged[0]; i++ {
			if record[i] == OPERATIONAL {
				operationalExists = true
				break
			}
		}
		if operationalExists || (i < len(record) && record[i] == DAMAGED) {
			return 0
		}

		newRecord := record[damaged[0]:]
		// make the next ? or . a .

		if len(newRecord) > 0 {
			newRecord = fmt.Sprintf("%c%s", OPERATIONAL, newRecord[1:])
		}

		return countArrangements(newRecord, damaged[1:])
	} else {
		left := fmt.Sprintf("%c%s", OPERATIONAL, record[1:])
		right := fmt.Sprintf("%c%s", DAMAGED, record[1:])
		l, r := countArrangements(left, damaged), countArrangements(right, damaged)
		return l + r
	}
}

func readInput(src string) ([]Record, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := []Record{}

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		record := lineSplit[0]
		damaged := []int{}

		damagesSplit := strings.Split(lineSplit[1], ",")
		for _, strNum := range damagesSplit {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return nil, err
			}
			damaged = append(damaged, num)
		}

		out = append(out, Record{
			record:  record,
			damaged: damaged,
		})
	}

	return out, nil
}
