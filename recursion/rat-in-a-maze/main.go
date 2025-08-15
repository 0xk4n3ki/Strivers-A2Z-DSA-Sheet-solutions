package main

import "fmt"

func main() {
	n := 4
	grid := [][]int{{1, 0, 0, 0}, {1, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 1, 1}}

	ans := maze1(n, grid)
	fmt.Printf("n: %d, ans: %v\n", n, ans)

	ans = maze2(n, grid)
	fmt.Printf("n: %d, ans: %v\n", n, ans)
}


// sol 2:
func maze2(n int, grid [][]int) []string {
	ans := []string {}
	vis := [][]int {}
	for i := 0; i < n; i++ {
		vis = append(vis, make([]int, n))
	}

	xVar := []int {1, 0, 0, -1}
	yVar := []int {0, -1, 1, 0}

	if grid[0][0] == 1 {
		backtracking2(0, 0, grid, n, &ans, "", &vis, xVar, yVar)
	}

	return ans
}

func backtracking2(i, j int, grid [][]int, n int, ans *[]string, path string, vis *[][]int, xVar, yVar []int) {
	if i == n-1 && j == n-1 {
		*ans = append(*ans, path)
		return
	}

	direction := "DLRU"

	for index := 0; index < 4; index++ {
		nexti := i + xVar[index]
		nextj := j + yVar[index]

		if nexti >= 0 && nextj >= 0 && nexti < n && nextj < n && 
		(*vis)[nexti][nextj] != 1 && grid[nexti][nextj] == 1 {
			(*vis)[i][j] = 1
			backtracking2(nexti, nextj, grid, n, ans, path+string(direction[index]), vis, xVar, yVar)
			(*vis)[i][j] = 0
		}
	}
}





// sol 1:
func maze1(n int, grid [][]int) []string {
	ans := []string {}
	vis := [][]int {}
	for i := 0; i < n; i++ {
		vis = append(vis, make([]int, n))
	}

	if grid[0][0] == 1 {
		backtracking1(0, 0, grid, n, &ans, "", &vis)
	}

	return ans
}

func backtracking1(i, j int, grid [][]int, n int, ans *[]string, path string, vis *[][]int) {
	if i == n-1 && j == n-1 {
		*ans = append(*ans, path)
		return
	}

	// downward
	if i + 1 < n && (*vis)[i+1][j] != 1 && grid[i+1][j] == 1 {
		(*vis)[i][j] = 1
		backtracking1(i+1, j, grid, n, ans, path+"D", vis)
		(*vis)[i][j] = 0
	}

	// left
	if j - 1 > 0 && (*vis)[i][j-1] != 1 && grid[i][j-1] == 1 {
		(*vis)[i][j] = 1
		backtracking1(i, j-1, grid, n, ans, path+"L", vis)
		(*vis)[i][j] = 0
	}

	// right
	if j + 1 < n && (*vis)[i][j+1] != 1 && grid[i][j+1] == 1 {
		(*vis)[i][j] = 1
		backtracking1(i, j+1, grid, n, ans, path+"R", vis)
		(*vis)[i][j] = 0
	}

	// upward
	if i - 1 > 0 && (*vis)[i-1][j] != 1 && grid[i-1][j] == 1 {
		(*vis)[i][j] = 1
		backtracking1(i-1, j, grid, n, ans, path+"U", vis)
		(*vis)[i][j] = 0
	}
}
