package main

import "fmt"

func main() {
	matrix := [][]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	bruteforceApp(matrix)

	optimalApp(matrix)
}


func optimalApp(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Printf("transpose: %v\n", matrix)
	
	for i := 0; i < len(matrix); i++ {
		reverseRow(matrix[i])
	}
	fmt.Printf("reverse: %v\n", matrix)
}

func reverseRow(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}


func bruteforceApp(matrix [][]int) {
	dup := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dup[i] = make([]int, len(matrix[i]))
	}
	
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dup[j][len(matrix)-i-1] = matrix[i][j]
		}
	}

	fmt.Printf("dup: %v\n", dup)
	fmt.Printf("mat: %v\n", matrix)
}