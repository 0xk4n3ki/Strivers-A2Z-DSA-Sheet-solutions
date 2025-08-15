package main

import (
	"fmt"
)

func main() {
	matrix := [][]int {{8, 6}, {10, 1}}

	bruteforce(matrix)
	optiBruteforce(matrix)
	optApp(matrix)

	matrix = [][]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	bruteforce(matrix)
	optiBruteforce(matrix)
	optApp(matrix)

	optiBruteforce([][]int {{10, 10}, {10, 9}})
}

func optApp(matrix [][]int) {
	m := len(matrix[0])
	low, high := 0, m-1

	for low <= high {
		mid := (low + high)/2

		row := maxColEle(matrix, mid)

		var left, right int

		if mid - 1 >= 0 {
			left = matrix[row][mid-1]
		}else {
			left = -1
		}

		if mid + 1 < m {
			right = matrix[row][mid+1]
		}else {
			right = -1
		}

		if matrix[row][mid] > left && matrix[row][mid] > right {
			fmt.Printf("matrix: %v, peak element index: (%d, %d)\n", matrix, row, mid)
			return
		}else if matrix[row][mid] < left {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("matrix: %v, peak element index: (%d, %d)\n", matrix, -1, -1)
}

func maxColEle(matrix [][]int, col int) int {
	maxEle, ind := -1, -1

	for i := 0; i < len(matrix); i++ {
		if matrix[i][col] > maxEle {
			maxEle = matrix[i][col]
			ind = i
		}
	}
	return ind
}

func optiBruteforce(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	row, col, ele := -1, -1, -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] > ele {
				ele = matrix[i][j]
				row = i
				col = j
			}
		}
	}
	fmt.Printf("matrix: %v, peak element index: (%d, %d)\n", matrix, row, col)
}

func bruteforce(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ele := matrix[i][j]
			
			if i - 1 >= 0 && ele <= matrix[i-1][j] {
				continue
			}

			if i + 1 < n && ele <= matrix[i+1][j] {
				continue
			}

			if j - 1 >= 0 && ele <= matrix[i][j-1] {
				continue
			}

			if j + 1 < m && ele <= matrix[i][j+1] {
				continue
			}

			fmt.Printf("matrix: %v, peak element index: (%d, %d)\n", matrix, i, j)
			return
		}
	}
	fmt.Printf("matrix: %v, peak element index: (%d, %d)\n", matrix, -1, -1)
}