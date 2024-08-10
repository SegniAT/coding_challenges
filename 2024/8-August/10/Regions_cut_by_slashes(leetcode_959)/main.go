package main

import "fmt"

func main() {
	//grid := []string{"/\\", "\\/"}
	//	grid := []string{"  ", "  "}
	grid := []string{" /", "/ "}
	res := regionsBySlashes(grid)
	fmt.Println(res)
}

func regionsBySlashes(grid []string) int {
	N := len(grid)
	regions := 0

	visited := make([][]bool, 3*N)
	for i := range visited {
		visited[i] = make([]bool, 3*N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '/' {
				visited[3*i][3*j+2] = true
				visited[3*i+1][3*j+1] = true
				visited[3*i+2][3*j] = true
			} else if grid[i][j] == '\\' {
				visited[3*i][3*j] = true
				visited[3*i+1][3*j+1] = true
				visited[3*i+2][3*j+2] = true
			}
		}
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= 3*N || y < 0 || y >= 3*N || visited[x][y] {
			return
		}
		visited[x][y] = true
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}

	for i := 0; i < 3*N; i++ {
		for j := 0; j < 3*N; j++ {
			if !visited[i][j] {
				dfs(i, j)
				regions++
			}
		}
	}

	return regions
}
