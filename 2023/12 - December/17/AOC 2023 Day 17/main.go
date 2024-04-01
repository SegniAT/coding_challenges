package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Direction rune

// const (
// 	UP    Direction = 'U'
// 	RIGHT Direction = 'R'
// 	DOWN  Direction = 'D'
// 	LEFT  Direction = 'L'
// )

// var dirMap map[Direction][]int = map[Direction][]int{
// 	UP:    {-1, 0},
// 	RIGHT: {0, 1},
// 	DOWN:  {1, 0},
// 	LEFT:  {0, -1},
// }

// type MemoEntry struct {
// 	minHeatLoss int
// 	prevDir     Direction
// }

// var memo map[string]MemoEntry

// func main() {
// 	input, err := loadInput("sample_input.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	rows, cols := len(input), len(input[0])

// 	// prettyPrint(input)

// 	minHeatLoss := make([][]int, rows)

// 	for rowInd, row := range input {
// 		minHeatLoss[rowInd] = make([]int, cols)
// 		for colInd := range row {
// 			minHeatLoss[rowInd][colInd] = math.MaxInt
// 		}
// 	}

// 	minHeatLoss[0][0] = 0
// 	walk(&input, []int{0, 1}, &minHeatLoss, RIGHT, 2, 0)
// 	walk(&input, []int{1, 0}, &minHeatLoss, DOWN, 2, 0)

// 	// fmt.Println("Min heat loss matrix: ")
// 	// prettyPrint(minHeatLoss)

// 	fmt.Println("Min heat loss to reach bottom right: ", minHeatLoss[rows-1][cols-1])
// }

// func walk(matrix *[][]int, coord []int, minHeatLoss *[][]int, dir Direction, count int, prev int) int {
// 	key := coordStr(coord)
// 	currRow, currCol := coord[0], coord[1]

// 	if val, ok := memo[key]; ok && val.minHeatLoss <= prev + (*matrix)[currRow][currCol] {
// 		return val.minHeatLoss
// 	}

// 	if count < 0 || !isValidCoord(matrix, coord) {
// 		return math.MaxInt
// 	}

// 	if _, ok := visited[coordStr(coord)]; ok {
// 		return
// 	} else {
// 		visited[coordStr(coord)] = true
// 	}

// 	currMin := (*minHeatLoss)[currRow][currCol]
// 	newMin := prev + (*matrix)[currRow][currCol]
// 	if newMin > currMin {
// 		return
// 	} else {
// 		(*minHeatLoss)[currRow][currCol] = newMin
// 	}

// 	newVisited := copyMap(visited)

// 	switch dir {
// 	case UP:
// 		// continue going UP, go LEFT or go RIGHT
// 		walk(matrix, addCoords(coord, dirMap[UP]), newVisited, minHeatLoss, UP, count-1, newMin)
// 		walk(matrix, addCoords(coord, dirMap[LEFT]), newVisited, minHeatLoss, LEFT, 2, newMin)
// 		walk(matrix, addCoords(coord, dirMap[RIGHT]), newVisited, minHeatLoss, RIGHT, 2, newMin)
// 	case RIGHT:
// 		// continue going RIGHT, go UP or go DOWN
// 		walk(matrix, addCoords(coord, dirMap[RIGHT]), newVisited, minHeatLoss, RIGHT, count-1, newMin)
// 		walk(matrix, addCoords(coord, dirMap[UP]), newVisited, minHeatLoss, UP, 2, newMin)
// 		walk(matrix, addCoords(coord, dirMap[DOWN]), newVisited, minHeatLoss, DOWN, 2, newMin)
// 	case DOWN:
// 		// continue going DOWN, go LEFT or go RIGHT
// 		walk(matrix, addCoords(coord, dirMap[DOWN]), newVisited, minHeatLoss, DOWN, count-1, newMin)
// 		walk(matrix, addCoords(coord, dirMap[LEFT]), newVisited, minHeatLoss, LEFT, 2, newMin)
// 		walk(matrix, addCoords(coord, dirMap[RIGHT]), newVisited, minHeatLoss, RIGHT, 2, newMin)
// 	case LEFT:
// 		// continue going UP, go LEFT or go RIGHT
// 		walk(matrix, addCoords(coord, dirMap[LEFT]), newVisited, minHeatLoss, LEFT, count-1, newMin)
// 		walk(matrix, addCoords(coord, dirMap[UP]), newVisited, minHeatLoss, UP, 2, newMin)
// 		walk(matrix, addCoords(coord, dirMap[DOWN]), newVisited, minHeatLoss, DOWN, 2, newMin)
// 	}

// }

// func copyMap(visited map[string]bool) map[string]bool {
// 	newMap := make(map[string]bool)

// 	for key, val := range visited {
// 		newMap[key] = val
// 	}
// 	return newMap
// }

// func addCoords(coord1, coord2 []int) []int {
// 	return []int{coord1[0] + coord2[0], coord1[1] + coord2[1]}
// }

// func coordStr(coord []int) string {
// 	return fmt.Sprintf("%d,%d", coord[0], coord[1])
// }

// func isValidCoord(matrix *[][]int, coord []int) bool {
// 	return (0 <= coord[0] && coord[0] < len(*matrix)) && (0 <= coord[1] && coord[1] < len((*matrix)[0]))
// }

// func prettyPrint(matrix [][]int) {
// 	fmt.Println(strings.Repeat("_", len(matrix[0])*2))
// 	for _, row := range matrix {
// 		fmt.Println(row)
// 	}
// 	fmt.Println(strings.Repeat("_", len(matrix[0])*2))
// }

// func loadInput(src string) ([][]int, error) {
// 	file, err := os.Open(src)
// 	if err != nil {
// 		return nil, err
// 	}

// 	out := [][]int{}
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), "")
// 		lineOut := []int{}

// 		for _, l := range line {
// 			num, err := strconv.Atoi(l)
// 			if err != nil {
// 				return nil, err
// 			}

// 			lineOut = append(lineOut, num)
// 		}

// 		out = append(out, lineOut)
// 	}

// 	return out, nil
// }
