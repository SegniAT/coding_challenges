package main

import (
	"fmt"
	"math"
)

func main() {
}

func findMinHeightTrees(n int, edges [][]int) []int {
	edgesMap := make(map[int][]int)
	for i := range n {
		edgesMap[i] = []int{}
	}

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		edgesMap[node1] = append(edgesMap[node1], node2)
		edgesMap[node2] = append(edgesMap[node2], node1)
	}

	heightForRoot := make(map[int]int)
	minHeight := math.MaxInt

	// for every root
	for root := range n {
		visited := make(map[int]bool)
		maxDepth := 0
		dfs(root, &maxDepth, 0, edgesMap, &visited)
		if maxDepth < minHeight {
			minHeight = maxDepth
		}
		heightForRoot[root] = maxDepth
	}

	fmt.Println(heightForRoot)

	res := []int{}
	for key, val := range heightForRoot {
		if val == minHeight {
			res = append(res, key)
		}
	}

	return res
}

func dfs(node int, maxDepth *int, depth int, edgesMap map[int][]int, visited *map[int]bool) {
	if depth > *maxDepth {
		*maxDepth = depth
	}

	(*visited)[node] = true

	for _, neighbourNode := range edgesMap[node] {
		if !(*visited)[neighbourNode] {
			dfs(neighbourNode, maxDepth, depth+1, edgesMap, visited)
		}
	}
}
