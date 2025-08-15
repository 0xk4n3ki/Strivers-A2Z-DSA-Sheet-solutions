package main

import (
	"fmt"
)

func main() {
	matrix := [][]int {{1, 1, 1}, {0, 0, 1}, {0, 0, 0}}
	bruteforceApp(matrix)
	optApp(matrix)

	matrix = [][]int {{0, 0}, {0, 0}}
	bruteforceApp(matrix)
	optApp(matrix)
}

func optApp(matrix [][]int) {
	maxOnes, rowIndex := 0, -1

	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		low, high := 0, cols - 1

		for low <= high {
			mid := (low + high)/2

			if matrix[i][mid] == 1 {
				high = mid - 1
			}else {
				low = mid + 1
			}
		}

		onesCount := cols - low

		if onesCount > maxOnes {
			maxOnes = onesCount
			rowIndex = i
		}
	}

	fmt.Printf("matrix: %v, row index: %d\n", matrix, rowIndex)
}

func bruteforceApp(matrix [][]int) {
	maxOnes := 0
	rowIndex := -1
	rows, col := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		count := 0
		for j:= 0; j < col; j++ {
			if matrix[i][j] == 0 {
				count++
			}else {
				break
			}
		}
		count = col - count
		if count > maxOnes {
			maxOnes = count
			rowIndex = i
		}
	}

	fmt.Printf("arr: %v, index: %d\n", matrix, rowIndex)
}