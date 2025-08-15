package main
import "fmt"

func main() {
	ans := nQueens(4)
	fmt.Printf("ans: %v\n", ans)
}

func nQueens(n int) [][]string {
	ans := [][]string {}
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		row := make([]rune, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}

	leftRow := make([]int, n)
	upperDiagonal := make([]int, 2*n-1)
	lowerDiagonal := make([]int, 2*n-1)

	solve2(0, &board, &ans, &leftRow, &upperDiagonal, &lowerDiagonal, n)

	return ans
}

func solve2(col int, board *[][]rune, ans *[][]string, leftRow, upperDiagonal, lowerDiagonal *[]int, n int) {
	if col == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string((*board)[i])
		}
		*ans = append(*ans, tmp)
		return
	}

	for row := 0; row < n; row++ {
		if (*leftRow)[row] == 0 && (*lowerDiagonal)[row+col] == 0 && (*upperDiagonal)[n-1+col-row] == 0 {
			(*board)[row][col] = 'Q'
			(*leftRow)[row] = 1
			(*lowerDiagonal)[row+col] = 1
			(*upperDiagonal)[n-1 + col-row] = 1

			solve2(col+1, board, ans, leftRow, upperDiagonal, lowerDiagonal, n)

			(*board)[row][col] = '.'
			(*leftRow)[row] = 0
			(*lowerDiagonal)[row+col] = 0
			(*upperDiagonal)[n-1 + col-row] = 0
		}
	}
}