package main

import "fmt"

func main() {
	matrix := [][]int {{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}};
	target := 8

	betterApp(matrix, target)
	optApp(matrix, target)

	betterApp(matrix, 34)
	optApp(matrix, 34)
}

func optApp(matrix [][]int, target int) {
	n, m := len(matrix), len(matrix[0])

	row, col := 0, m-1

	for row < n && col >= 0 {
		if matrix[row][col] == target {
			fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, true)
			return
		}else if matrix[row][col] > target {
			col--
		}else {
			row++
		}
	}
	fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, false)
}

func betterApp(matrix [][]int, target int) {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		low, high := 0, m-1
		
		if target >= matrix[i][low] && target <= matrix[i][high] {
			for low <= high {
				mid := (low + high)/2

				if matrix[i][mid] == target {
					fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, true)
					return
				}else if target < matrix[i][mid] {
					high = mid - 1
				}else {
					low = mid + 1
				}
			}
		}
	}
	fmt.Printf("matrix: %v, target: %d, ans: %t\n", matrix, target, false)
}