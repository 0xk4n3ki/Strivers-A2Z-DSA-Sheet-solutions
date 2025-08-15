package main

import "fmt"

func main() {
	matrix := [][]int {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	spiral(matrix)
}

func spiral(matrix [][]int) {
	topRow, bottomRow, leftCol, rightCol := 0, len(matrix)-1, 0, len(matrix[0])-1

	fmt.Printf("matrix: %v\n", matrix)
	fmt.Print("spiral: ")
	for topRow <= bottomRow && leftCol <= rightCol {
		for i := leftCol; i <= rightCol; i++ {
			fmt.Printf("%d ", matrix[topRow][i])
		}
		topRow++

		for i := topRow; i <= bottomRow; i++ {
			fmt.Printf("%d ", matrix[i][rightCol])
		}
		rightCol--

		if topRow <= bottomRow {
			for i := rightCol; i >= leftCol; i-- {
				fmt.Printf("%d ", matrix[bottomRow][i])
			}
			bottomRow--
		}

		if leftCol <= rightCol {
			for i := bottomRow; i >= topRow; i-- {
				fmt.Printf("%d ", matrix[i][leftCol])
			}
			leftCol++
		}
	}
	fmt.Println()
}