package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	vis := make([]int, n+1)
	start := 1
	ans := []int {}

	dfs(start, adj, &vis, &ans)
	fmt.Printf("ans: %v\n", ans)
}

func dfs(start int, adj [][]int, vis , ans *[]int) {
	(*vis)[start] = 1
	*ans = append(*ans, start)

	for _, i := range adj[start] {
		if (*vis)[i] != 1 {
			dfs(i, adj, vis, ans)
		}
	}
}