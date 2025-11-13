package main

import (
	"fmt"
)

func main() {
	v := 4
	edges := [][]int {{0, 1}, {1, 2}}
	fmt.Printf("v: %d, edges: %v, ans: %d\n", v, edges, calc(v, edges))

	v = 7
	edges = [][]int {{0, 1}, {1, 2}, {2, 3}, {4, 5}}
	fmt.Printf("v: %d, edges: %v, ans: %d\n", v, edges, calc(v, edges))
}

func calc(v int, edges [][]int) int {
	ans := 0
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		tmp := make([]int, v)
		adj[i] = tmp
	}

	for _, i := range edges {
		adj[i[0]][i[1]] = 1
		adj[i[1]][i[0]] = 1
	}

	vis := make([]int, v)

	for i := 0; i < v; i++ {
		if vis[i] != 1 {
			ans++
			dfs(adj, &vis, i)
		}
	}
	
	return ans
}

func dfs(adj [][]int, vis *[]int, i int) {
	(*vis)[i] = 1

	for j := 0; j < len(adj); j++ {
		if adj[i][j] == 1 && (*vis)[j] != 1 {
			dfs(adj, vis, j)
		}
	}
}