package main

import (
	"fmt"
)

func main() {
	n := 4
	ans := queens1(n)
	fmt.Printf("n: %d, ans: %v\n", n, ans)

	n = 4
	ans = queens2(n)
	fmt.Printf("n: %d, ans: %v\n", n, ans)
}


// approach 2:
func queens2(n int) [][]string {
	ans := [][]string {}
	sub := [][]rune {}
	for i := 0; i < n; i++ {
		tmp := make([]rune, n)
		for j := 0; j < n; j++ {
			tmp[j] = '.'
		}
		sub = append(sub,tmp)
	}

	leftrow, upperDiagonal, lowerDiagonal := make([]int, n), make([]int, 2*n-1), make([]int, 2*n-1)
	solve2(0, &sub, &ans, &leftrow, &upperDiagonal, &lowerDiagonal, n)
	
	return ans
}

func solve2(col int, sub *[][]rune, ans *[][]string, leftrow, upperDiagonal, lowerDiagonal *[]int, n int) {
	if col == n {
		tmp := []string {}
		for i := 0; i < n; i++ {
			tmp = append(tmp, string((*sub)[i]))
		}
		*ans = append(*ans, tmp)
		return
	}

	for row := 0; row < n; row++ {
		if (*leftrow)[row] == 0 && (*lowerDiagonal)[n-1 + col-row] == 0 && (*upperDiagonal)[row+col] == 0 {
			(*sub)[row][col] = 'Q'
			(*leftrow)[row] = 1
			(*upperDiagonal)[row+col] = 1
			(*lowerDiagonal)[n-1 + col-row] = 1

			solve2(col+1, sub, ans, leftrow, upperDiagonal, lowerDiagonal, n)

			(*sub)[row][col] = '.'
			(*leftrow)[row] = 0
			(*upperDiagonal)[row+col] = 0
			(*lowerDiagonal)[n-1 + col-row] = 0
		}
	}
}








// approach 1:
func queens1(n int) [][]string {
	ans := [][]string {}
	sub := [][]rune {}
	
	for i := 0; i < n; i++ {
		tmp := []rune {}
		for j := 0; j < n; j++ {
			tmp = append(tmp, '.')
		}
		sub = append(sub, tmp)
	}

	backtracking1(0, &sub, &ans, n)

	return ans
}

func backtracking1(col int, sub *[][]rune, ans *[][]string, n int) {
	if col == n {
		tmp := []string {}
		for i := 0; i < n; i++ {
			tmp = append(tmp, string((*sub)[i]))
		}
		*ans = append(*ans, tmp)
		return
	}

	for row := 0; row < n; row++ {
		if isSafe1(row, col, sub, n) {
			(*sub)[row][col] = 'Q'
			backtracking1(col+1, sub, ans, n)
			(*sub)[row][col] = '.'
		}
	}
}

func isSafe1(row, col int, sub *[][]rune, n int) bool {
	dupRow, dupCol := row, col

	// upper left diagonal
	for row >= 0 && col >= 0 {
		if (*sub)[row][col] == 'Q' {
			return false
		}
		row--
		col--
	}

	col, row = dupCol, dupRow
	// row
	for col >= 0 {
		if (*sub)[row][col] == 'Q' {
			return false
		}
		col--
	}

	col, row = dupCol, dupRow
	// lower left diagonal
	for row < n && col >= 0 {
		if (*sub)[row][col] == 'Q' {
			return false
		}
		row++
		col--
	}

	return true
}