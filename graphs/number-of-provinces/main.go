package main

import (
	"fmt"
)

func main() {
	adj := [][]int {{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}}
	fmt.Printf("adj: %v, ans: %d\n", adj, calc(adj))

	adj = [][]int {{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	fmt.Printf("adj: %v, ans: %d\n", adj, calc(adj))

	adj = [][]int {{1, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 1, 0}, {0, 0, 0, 1}}
	fmt.Printf("adj: %v, ans: %d\n", adj, calc(adj))
}

func calc(adj [][]int) int {
	ans := 0
	n := len(adj)
	arr := [][]int {}

	vis := make([]int, n)

	for i := 0; i < n; i++ {
		if vis[i] != 1 {
			ans++
			tmp := []int {}
			dfs(adj, &vis, i, &tmp)
			arr = append(arr, tmp)
		}
	}
	fmt.Printf("arr: %v\n", arr)
	return ans
}

func dfs(adj [][]int, vis *[]int, i int, tmp *[]int) {
	(*vis)[i] = 1
	*tmp = append(*tmp, i)

	for j := 0; j < len(adj); j++ {
		if adj[i][j] == 1 && (*vis)[j] != 1 {
			dfs(adj, vis, j, tmp)
		}
	}
}