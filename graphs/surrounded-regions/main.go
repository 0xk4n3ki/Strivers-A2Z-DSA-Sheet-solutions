package main

import (
	"fmt"
)

func main() {
	matrix := [][]int {{1, 1, 1, 1}, {1, 0, 1, 1}, {1, 0, 0, 1} ,{1, 0, 1, 1}, {1, 1, 0, 0}}
	fmt.Printf("ans: %v\n", calc(matrix))

	matrix = [][]int {{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}}
	fmt.Printf("ans: %v\n", calc(matrix))
}

func calc(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	vis, ans := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
		ans[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		if matrix[0][j] == 0 && vis[0][j] != 1 {
			dfs(matrix, 0, j, vis)
		}

		if matrix[n-1][j] == 0 && vis[n-1][j] != 1 {
			dfs(matrix, n-1, j, vis)
		}
	}
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 && vis[i][0] != 1 {
			dfs(matrix, i, 0, vis)
		}
		if matrix[i][m-1] == 0 && vis[i][m-1] != 1 {
			dfs(matrix, i, m-1, vis)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				ans[i][j] = 1
			}
			if vis[i][j] != 1 && matrix[i][j] == 0 {
				ans[i][j] = 1
			}
		}
	}
	return ans
}

func dfs(matrix [][]int, row, col int, vis [][]int) {
	vis[row][col] = 1
	n, m := len(matrix), len(matrix[0])

	drow, dcol := []int {-1, 0, +1, 0}, []int {0, +1, 0, -1}
	for i := 0; i < 4; i++ {
		nrow, ncol := row+drow[i], col+dcol[i]

		if nrow >= 0 && ncol >= 0 && nrow < n && ncol < m && matrix[nrow][ncol] == 0 && vis[nrow][ncol] != 1 {
			dfs(matrix, nrow, ncol, vis)
		}
	}
}