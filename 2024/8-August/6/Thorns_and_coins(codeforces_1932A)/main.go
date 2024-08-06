package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	EMPTY_CELL = '.'
	COIN_CELL  = '@'
	THORN_CELL = '*'
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	tests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Atoi: error")
		panic(err)
	}

	for range tests {
		scanner.Scan()
		scanner.Scan()
		path := scanner.Text()
		cache := make(map[string]int)
		fmt.Println(helper(path, 0, 0, &cache))
	}
}

func helper(path string, sum, ind int, cache *map[string]int) int {
	N := len(path)
	if ind >= N {
		return sum
	}

	key := fmt.Sprintf("%d,%d", ind, sum)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	if path[ind] == COIN_CELL {
		sum++
	} else if path[ind] == THORN_CELL {
		return sum
	}

	opt1 := helper(path, sum, ind+1, cache)
	opt2 := helper(path, sum, ind+2, cache)

	(*cache)[key] = max(opt1, opt2)

	return (*cache)[key]
}
