package main

import "fmt"


func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// adjacencyMatrix(n, m)
	// adjacencyList(n, m)
	// weightedMatrix(n, m)
	weightedList(n, m)
}

type Pair [2]int
func weightedList(n, m int) {
	adj := make([][]Pair, n+1)
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)

		adj[u] = append(adj[u], [2]int {v, w})
		adj[v] = append(adj[v], [2]int {u, w})
	}
	
	for i := 0; i < m; i++ {
		fmt.Printf("%d: %v\n", i, adj[i])
	}
}



func weightedMatrix(n, m int) {
	adj := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		adj[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		var w, u, v int
		fmt.Scan(&u, &v, &w)
		adj[u][v] = w
		adj[v][u] = w
	}

	for i := 0; i < n+1; i++ {
		fmt.Printf("%v\n", adj[i])
	}
}



func adjacencyList(n, m int) {
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u) // for directed graph, simply remove this line
	}

	for i := 0; i < n+1; i++ {
		fmt.Printf("%d: %v\n", i, adj[i])
	}
}

func adjacencyMatrix(n, m int) {
	adj := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		adj[i] = make([]int, n+1)
	}
	
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		adj[u][v] = 1
		adj[v][u] = 1 // for directed graph, simply remove this line
	}
	
	for i := 0; i < n+1; i++ {
		fmt.Printf("%v\n", adj[i])
	}
}