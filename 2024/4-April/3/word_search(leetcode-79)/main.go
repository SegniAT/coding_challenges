package main

func main() {
}

var DIRS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up right down left

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			visited := genVisited(rows, cols)
			res := helper(board, visited, []int{row, col}, word, 0)
			if res {
				return true
			}
		}
	}

	return false
}

func helper(board [][]byte, visited [][]bool, cell []int, word string, wordInd int) bool {
	row, col := cell[0], cell[1]

	boardChar := board[row][col]
	if word[wordInd] != boardChar {
		return false
	} else if wordInd == len(word)-1 {
		return true
	}

	visited[row][col] = true

	res := false
	for _, dir := range DIRS {
		newCell := []int{row + dir[0], col + dir[1]}
		newRow, newCol := newCell[0], newCell[1]
		if inRange(board, newCell) && !visited[newRow][newCol] {
			res = res || helper(board, visited, newCell, word, wordInd+1)
		}
	}

	visited[row][col] = false

	return res
}

func inRange(board [][]byte, cell []int) bool {
	row, col := cell[0], cell[1]
	return (0 <= row && row < len(board)) && (0 <= col && col < len(board[0]))
}

func genVisited(rows, cols int) [][]bool {
	visited := make([][]bool, rows)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	return visited
}
