package main

import (
	"fmt"
)

func main() {
	grid := [][]int {{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	fmt.Printf("ans: %d\n", calc(grid))

	grid = [][]int {{0, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {0, 1, 1, 0}}
	fmt.Printf("ans: %d\n", calc(grid))

	grid = [][]int {{0, 0, 0, 1, 1}, {0, 0, 1, 1, 0}, {0, 1, 0, 0, 0}, {0, 1, 1, 0, 0}, {0, 0, 0, 1, 1}}
	fmt.Printf("ans: %d\n", calc(grid))
}

func calc(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		if grid[0][i] == 1 && vis[0][i] != 1 {
			dfs(grid, 0, i, vis)
		}
		if grid[n-1][i] == 1 && vis[n-1][i] != 1 {
			dfs(grid, n-1, i, vis)
		}
	}

	for i := 0; i < n; i++ {
		if grid[i][0] == 1 && vis[i][0] != 1 {
			dfs(grid, i, 0, vis)
		}
		if grid[i][m-1] == 1 && vis[i][m-1] != 1 {
			dfs(grid, i, m-1, vis)
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && vis[i][j] != 1 {
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, x, y int, vis [][]int) {
	n, m := len(grid), len(grid[0])
	vis[x][y] = 1

	drow, dcol := []int {-1, 0, +1, 0}, []int {0, +1, 0, -1}

	for i := 0; i < 4; i++ {
		nrow, ncol := x+drow[i], y+dcol[i]
		if nrow >= 0 && ncol >= 0 && nrow < n && ncol < m && grid[nrow][ncol] == 1 &&
			vis[nrow][ncol] != 1 {
				dfs(grid, nrow, ncol, vis)
			}
	}
}