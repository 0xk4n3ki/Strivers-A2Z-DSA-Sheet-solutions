package main

import (
	"fmt"
)

func main() {
	ans := queen(4)
	fmt.Printf("ans: %v\n", ans)
}

func queen(n int) [][]string {
	ans := [][]string {}
	base := make([][]rune, n)
	for i := 0; i < n; i++ {
		arr := make([]rune, n)
		for j := 0; j < n; j++ {
			arr[j] = '.'
		}
		base[i] = arr
	}

	solve(0, &base, &ans, n)
	return  ans
}

func solve(col int, base *[][]rune, ans *[][]string, n int) {
	if col == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string((*base)[i])
		}
		*ans = append(*ans, tmp)
		return
	}

	for row := 0; row < n; row++{
		if isSafe(row, col, base, n) {
			(*base)[row][col] = 'Q'
			solve(col+1, base, ans, n)
			(*base)[row][col] = '.'
		}
	}
}

func isSafe(row, col int, base *[][]rune, n int) bool {
	dupRow, dupCol := row, col
	for row >= 0 && col >= 0 {
		if (*base)[row][col] == 'Q' {
			return false
		}
		row--
		col--
	}

	row, col = dupRow, dupCol
	for col >= 0 {
		if (*base)[row][col] == 'Q' {
			return false
		}
		col--
	}

	row, col = dupRow, dupCol
	for row < n && col >= 0 {
		if (*base)[row][col] == 'Q' {
			return false
		}
		row++
		col--
	}

	return true
}
