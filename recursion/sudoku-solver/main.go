package main

import (
	"fmt"
)

func main() {
	board := [][]rune {
        {'9', '5', '7', '.', '1', '3', '.', '8', '4'},
        {'4', '8', '3', '.', '5', '7', '1', '.', '6'},
        {'.', '1', '2', '.', '4', '9', '5', '3', '7'},
        {'1', '7', '.', '3', '.', '4', '9', '.', '2'},
        {'5', '.', '4', '9', '7', '.', '3', '6', '.'},
        {'3', '.', '9', '5', '.', '8', '7', '.', '1'},
        {'8', '4', '5', '7', '9', '.', '6', '1', '3'},
        {'.', '9', '1', '.', '3', '6', '.', '7', '5'},
        {'7', '.', '6', '1', '8', '5', '4', '.', '9'},
    }

	solve(&board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func solve(board *[][]rune) bool {
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			if (*board)[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					
					if isValid(*board, i, j, c) {

						(*board)[i][j] = c

						if solve(board) {
							return true
						}else {
							(*board)[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]rune, row, col int, c rune) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == c {
			return false
		}

		if board[row][col] == c {
			return false
		}

		if board[3 * (row/3) + i/3][3 * (col/3) + i%3] == c {
			return false
		}
	}
	return true
}