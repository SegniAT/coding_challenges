package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	SAFE = '0'
	TRAP = '1'
)

/*
NOTE: The smart solution: just check if two TRAPs appear vertically next to each other, that's the blocker
*/

/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for range tests {
		path := make([]string, 2)

		scanner.Scan()
		cols, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		scanner.Scan()
		path[0] = scanner.Text()
		scanner.Scan()
		path[1] = scanner.Text()

		var col int
		for col = range cols {
			if path[0][col] == TRAP && path[0][col] == path[1][col] {
				fmt.Println("NO")
				break
			}
		}

		if col == cols-1 {
			fmt.Println("YES")
		}
	}
}
*/

// N, NE, E, SE, S
var DIRS = [][2]int{
	{-1, 0}, {-1, 1},
	{0, 1}, {1, 1},
	{1, 0},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < tests; i++ {
		scanner.Scan()
		cols, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		path := make([]string, 2)
		scanner.Scan()
		path[0] = scanner.Text()
		scanner.Scan()
		path[1] = scanner.Text()

		// Use a cache to store the state of each cell
		cache := make([][]int, 2)
		for j := 0; j < 2; j++ {
			cache[j] = make([]int, cols)
			for k := 0; k < cols; k++ {
				cache[j][k] = -1 // -1 means unvisited
			}
		}

		// Start from (0, 0)
		if dfs(path, cols, 0, 0, cache) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func dfs(path []string, cols, row, col int, cache [][]int) bool {
	// Out of bounds or trap cell
	if row < 0 || row >= 2 || col < 0 || col >= cols || path[row][col] == TRAP {
		return false
	}

	// Reached the destination
	if row == 1 && col == cols-1 {
		return true
	}

	// If already visited, return the cached result
	if cache[row][col] != -1 {
		return cache[row][col] == 1
	}

	// Mark the current cell as visited
	cache[row][col] = 0

	for _, dir := range DIRS {
		if dfs(path, cols, row+dir[0], col+dir[1], cache) {
			cache[row][col] = 1
			return true
		}
	}

	return false
}
