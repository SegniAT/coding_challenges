package main

import "fmt"

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},

		{'.', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},

		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}

	fmt.Println(isValidSudokuSmarter(board))
}

func isValidSudokuSmarter(board [][]byte) bool {

	hashMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			value := string(board[i][j])
			if value == "." {
				continue
			}

			_, rowCheck := hashMap[value+" in row "+string(rune(i))]
			_, colCheck := hashMap[value+" in col "+string(rune(j))]
			_, gridCheck := hashMap[value+" in grid ("+string(rune(+i/3))+","+string(rune(j/3))+")"]

			if rowCheck || colCheck || gridCheck {
				return false
			} else {
				hashMap[value+" in row "+string(rune(i))] = true
				hashMap[value+" in col "+string(rune(j))] = true
				hashMap[value+" in grid ("+string(rune(+i/3))+","+string(rune(j/3))+")"] = true
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	if !checkRowsHealth(board) || !checkColumnsHealth(board) {
		return false
	}

	// check sub-boxes
	origins := [][]int{{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6}}
	for _, origin := range origins {
		subBoxNums := make(map[byte]bool)

		for row := origin[0]; row < origin[0]+3; row++ {
			for col := origin[1]; col < origin[1]+3; col++ {
				currChar := board[row][col]
				if currChar == '.' {
					continue
				}

				if _, ok := subBoxNums[currChar]; ok {
					return false
				} else {
					subBoxNums[currChar] = true
				}
			}
		}
	}

	return true
}

func checkRowsHealth(board [][]byte) bool {
	rowNums := make(map[byte]bool)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			currChar := board[row][col]
			if currChar == '.' {
				continue
			}

			if _, ok := rowNums[currChar]; ok {
				return false
			} else {
				rowNums[currChar] = true
			}
		}
		rowNums = make(map[byte]bool)
	}
	return true
}

func checkColumnsHealth(board [][]byte) bool {
	colNums := make(map[byte]bool)
	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			currChar := board[row][col]
			if currChar == '.' {
				continue
			}

			if _, ok := colNums[currChar]; ok {
				return false
			} else {
				colNums[currChar] = true
			}
		}
		colNums = make(map[byte]bool)
	}
	return true
}
