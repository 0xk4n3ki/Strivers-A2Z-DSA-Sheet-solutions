package main

import (
	"fmt"
)

func main() {
	matrix := [][]int {{1, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 1, 1}}
	fmt.Printf("matrix: %v, size: %d\n", matrix, len(matrix[1]))

	bruteForceApp(matrix)
	inPlaceBruteForce(matrix)

	matrix = [][]int {{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	betterApp(matrix)

	matrix = [][]int {{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	optimalApp(matrix)
}

func optimalApp(matrix [][]int) {
	col0 := 1

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0

				if j != 0 {
					matrix[0][j] = 0
				}else {
					col0 = 0
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0|| matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if col0 == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}


	fmt.Printf("optimal approach by inplace: %v\n", matrix)
}


func betterApp(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}

	fmt.Printf("betterApp: %v\n", matrix)
}


func bruteForceApp(matrix [][]int) {
	dup := make([][]int, len(matrix))

	for i:= range matrix {
		dup[i] = make([]int, len(matrix[i]))
		copy(dup[i], matrix[i])
	}

	fmt.Printf("before: %v\n", dup)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for x := 0; x < len(matrix); x++ {
					for y := 0; y < len(matrix); y++ {
						if x == i || y == j {
							dup[x][y] = 0
						}
					}
				}
			}
		}
	}
	fmt.Printf("bruteforce: %v\n", dup)
}

func inPlaceBruteForce(mat [][]int) {
	matrix := make([][]int, len(mat))
	for i := range mat {
		matrix[i] = make([]int, len(mat[i]))
		copy(matrix[i], mat[i])
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix[i]); k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = -1
					}
				}
				for k := 0; k < len(matrix); k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = -1
					}
				}
			}
		}

	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Printf("inplace bruteforce: %v\n", matrix)
}