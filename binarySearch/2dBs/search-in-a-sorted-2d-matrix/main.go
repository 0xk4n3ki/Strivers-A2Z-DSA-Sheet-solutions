package main

import (
	"fmt"
)

func main() {
	matrix := [][]int {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	target := 8

	betterApp(matrix, target)
	optimalApp(matrix, target)

	betterApp([][]int {{1, 2, 4}, {6, 7, 8}, {9, 10, 34}}, 78)
	optimalApp([][]int {{1, 2, 4}, {6, 7, 8}, {9, 10, 34}}, 78)
}

func optimalApp(matrix [][]int, target int) {
	n, m := len(matrix), len(matrix[0])

	low, high := 0, n*m - 1

	for low <= high {
		mid := (low + high)/2

		row, col := mid/m, mid % m

		if matrix[row][col] == target {
			fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, true)
			return
		}else if matrix[row][col] > target {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, false)
}

func betterApp(matrix [][]int, target int) {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		if matrix[i][0] <= target && target <= matrix[i][m-1] {
			low, high := 0, m-1

			for low <= high {
				mid := (low + high)/2

				if matrix[i][mid] == target {
					fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, true)
					return
				}else if matrix[i][mid] > target {
					high = mid - 1
				}else {
					low = mid + 1
				}
			}
		}
	}
	fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, false)
}