package main

import (
	"fmt"
)

func main() {
	M := [][]int {
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 1, 1, 0},
	}
	ans := brute(M)
	fmt.Printf("ans: %d\n", ans)

	ans = optimal(M)
	fmt.Printf("ans: %d\n", ans)
}

func optimal(mat [][]int) int {
	top, bottom := 0, len(mat)-1

	for top < bottom {
		if mat[top][bottom] == 1 {
			top++
		}else {
			bottom--
		}
	}
	if top > bottom {
		return -1
	}

	for i := 0; i < len(mat); i++ {
		if i == top {
			continue
		}
		if mat[top][i] == 0 && mat[i][top] == 1 {
			continue
		}else {
			return -1
		}
	}
	return top
}

func brute(mat [][]int) int {
	n := len(mat)
	
	for i := 0; i < n; i++ {
		rowSum, colSum := 0, 0

		for j := 0; j < n; j++ {
			rowSum += mat[i][j]
			colSum += mat[j][i]
		}

		if rowSum == 0 && colSum == n-1 {
			return i
		}
	}
	return -1
}