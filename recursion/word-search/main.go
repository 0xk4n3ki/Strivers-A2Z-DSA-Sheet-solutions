package main

import "fmt"

func main() {
	word := "ABCCED"
	grid := [][]byte {{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	ans := search(word, grid)
	fmt.Printf("word: %s, ans: %t\n", word, ans)

	word = "ABCB"
	ans = search(word, grid)
	fmt.Printf("word: %s, ans: %t\n", word, ans)

	grid = [][]byte {{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	word = "SEE"
	ans = search(word, grid)
	fmt.Printf("word: %s, ans: %t\n", word, ans)
}

func search(word string, grid [][]byte) bool {
	m, n := len(grid), len(grid[0])

	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == word[index] {
				if backtracking(&grid, word, i, j, index, m, n) {
					return true
				}
			}
		}
	}
	return false
}

func backtracking(grid *[][]byte, word string, row, col, index, m, n int) bool {
	if index == len(word) {
		return true
	}

	if (row < 0 || col < 0 || row == m || col == n || (*grid)[row][col] != word[index] || (*grid)[row][col] == '!') {
		return false
	}

	c := (*grid)[row][col]
	(*grid)[row][col] = '!'

	top := backtracking(grid, word, row-1, col, index+1, m, n)
	right := backtracking(grid, word, row, col+1, index+1, m, n)
	bottom := backtracking(grid, word, row+1, col, index+1, m, n)
	left := backtracking(grid, word, row, col-1, index+1, m, n)

	(*grid)[row][col] = c

	return top || right || bottom || left
}