package main

func main() {

}

const (
	LIVE int = 0
	DEAD int = 0
)

func gameOfLife(board [][]int) {

	rows, cols := len(board), len(board[0])

	// N, NE, E, SE, S, SW, W, NW
	dirs := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	res := make([][]int, rows)
	for i := range board {
		res[i] = make([]int, cols)
		copy(res[i], board[i])
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeighbours, deadNeighbours := 0, 0

			// check neighbours
			for _, dir := range dirs {
				neighbour := []int{row + dir[0], col + dir[1]}
				if isValidCoordinate(neighbour, rows, cols) {
					if board[neighbour[0]][neighbour[1]] == 1 {
						liveNeighbours++
					} else {
						deadNeighbours++
					}
				}
			}

			if board[row][col] == 1 {
				if liveNeighbours < 2 {
					res[row][col] = 0
				} else if liveNeighbours < 4 {
					continue
				} else if liveNeighbours > 3 {
					res[row][col] = 0
				}
			} else {
				if liveNeighbours == 3 {
					res[row][col] = 1
				}
			}

		}
	}

	for i := range res {
		copy(board[i], res[i])
	}
}

func isValidCoordinate(coord []int, rows, cols int) bool {
	xCoord, yCoord := coord[0], coord[1]
	return (xCoord > -1 && xCoord < rows) && (yCoord > -1 && yCoord < cols)
}
