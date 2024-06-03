package main

import "fmt"

func main() {
	ans := validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5)
	fmt.Println(ans)
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	adjList := buildAdjacencyList(n, edges)
	visited := make(map[string]bool)
	ans := false
	dfs(adjList, &visited, source, destination, &ans)
	return ans
}

func dfs(adjList map[int][]int, visited *map[string]bool, curr int, destination int, ans *bool) {
	if *ans {
		return
	}

	if curr == destination {
		*ans = true
		return
	}

	adjNodes := adjList[curr]
	for _, node := range adjNodes {
		edge := fmt.Sprintf("%d-%d", curr, node)
		if _, ok := (*visited)[edge]; !ok {
			(*visited)[edge] = true
			dfs(adjList, visited, node, destination, ans)
		}
	}

}

func buildAdjacencyList(edgeCount int, edges [][]int) map[int][]int {
	adjList := make(map[int][]int)
	for i := range edgeCount {
		adjList[i] = []int{}
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	return adjList
}
