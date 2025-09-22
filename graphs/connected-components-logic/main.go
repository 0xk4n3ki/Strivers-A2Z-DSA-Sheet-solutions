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

	for i := 0; i < m; i++ {
		fmt.Printf("%d: %v\n", i, adj[i])
	}

	vis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if vis[i] == 0 {
			// fmt.Printf("%d ", i)
			vis[i] = 1
		}
	} 
	fmt.Printf("vis: %v\n", vis)
}