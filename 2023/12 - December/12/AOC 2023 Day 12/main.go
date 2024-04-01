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

	memo := make(map[string]int)

	sum := 0
	for _, inp := range input {
		// fmt.Println(inp.record, inp.damaged)
		multiply(&inp)
		arrangements := countArrangements(inp.record, inp.damaged, &memo)
		sum += arrangements
		// fmt.Println(arrangements)
		// fmt.Println()
	}
	fmt.Println("SUM: ", sum)

	// PART I
	// Attempt 1: 8073 (too high)
	// Attempt 2: 7633

	// PART II
	// Attempt 1: 22557823541987 (too low)
	// Attempt 2: 22031404990199
}

func multiply(record *Record) {
	initRecord, initDamaged := record.record, record.damaged
	for i := 0; i < 4; i++ {
		record.record = fmt.Sprintf("%s%c%s", record.record, '?', initRecord)
		record.damaged = append(record.damaged, initDamaged...)
	}
}

func countArrangements(record string, damaged []int, memo *map[string]int) int {
	if res, ok := (*memo)[fmt.Sprintf("%s-%d", record, len(damaged))]; ok {
		return res
	}

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
			(*memo)[fmt.Sprintf("%s-%d", record, len(damaged))] = 0
			return 0
		} else {
			(*memo)[fmt.Sprintf("%s-%d", record, len(damaged))] = 1
			return 1
		}
	}

	if record[0] == OPERATIONAL {
		ind := 0
		for ind < len(record) && record[ind] == OPERATIONAL {
			ind++
		}
		res := countArrangements(record[ind:], damaged, memo)
		(*memo)[fmt.Sprintf("%s-%d", record, len(damaged))] = res
		return res
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
			(*memo)[fmt.Sprintf("%s-%d", record, len(damaged))] = 0
			return 0
		}

		newRecord := record[damaged[0]:]
		// make the next ? or . a .

		if len(newRecord) > 0 {
			newRecord = fmt.Sprintf("%c%s", OPERATIONAL, newRecord[1:])
		}

		res := countArrangements(newRecord, damaged[1:], memo)
		(*memo)[fmt.Sprintf("%s-%d", record, len(damaged))] = res
		return res
	} else {
		left := fmt.Sprintf("%c%s", OPERATIONAL, record[1:])
		right := fmt.Sprintf("%c%s", DAMAGED, record[1:])
		l, r := countArrangements(left, damaged, memo), countArrangements(right, damaged, memo)
		(*memo)[fmt.Sprintf("%s-%d", record, len(damaged))] = l + r
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
