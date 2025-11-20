package main

import (
	"fmt"
)

func main() {
	v, e := 8, 7
	edges := [][]int{{1, 2}, {1, 3}, {2, 5}, {2, 6}, {3, 4}, {3, 7}, {7, 8}}
	fmt.Printf("ans: %t\n", check(edges, v, e))

	v, e = 8, 6
	edges = [][]int{{1, 2}, {2, 3}, {4, 5}, {5, 6}, {4, 6}, {7, 8}}
	fmt.Printf("ans: %t\n", check(edges, v, e))
}

func check(edges [][]int, v, e int) bool {
	vis := make([]int, v+1)
	adj := make([][]int, v+1)
	for _, i := range edges {
		u, v := i[0], i[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	fmt.Printf("adj: %v\n", adj)

	for i := 1; i <= v; i++ {
		if vis[i] != 1 {
			if dfs(adj, i, -1, vis) {
				return true
			}
		}
	}
	return false
}

func dfs(adj [][]int, i, parent int, vis []int) bool {
	vis[i] = 1

	for _, j := range adj[i] {
		if vis[j] != 1 {
			if dfs(adj, j, i, vis) {
				return true
			}
		} else if j != parent {
			return true
		}
	}
	return false
}
