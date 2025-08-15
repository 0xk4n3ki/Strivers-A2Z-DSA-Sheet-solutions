package main

import "fmt"

func main() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	ans := findPath(maze, len(maze))
	fmt.Printf("ans: %v\n", ans)
}

func findPath(maze [][]int, n int) []string {
	ans := []string{}
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}

	path := []rune{}

	if maze[0][0] == 1 {
		solve(0, 0, maze, n, &ans, &path, &visited)
	}
	return ans
}

func solve(i, j int, maze [][]int, n int, ans *[]string, path *[]rune, visited *[][]int) {
	if i == n-1 && j == n-1 {
		*ans = append(*ans, string(*path))
		return
	}

	// downward
	if i+1 < n && (*visited)[i+1][j] != 1 && maze[i+1][j] == 1 {
		(*visited)[i][j] = 1
		*path = append(*path, 'D')
		solve(i+1, j, maze, n, ans, path, visited)
		(*visited)[i][j] = 0
		*path = (*path)[:len(*path)-1]
	}

	// left
	if j-1 >= 0 && (*visited)[i][j-1] != 1 && maze[i][j-1] == 1 {
		(*visited)[i][j] = 1
		*path = append(*path, 'L')
		solve(i, j-1, maze, n, ans, path, visited)
		(*visited)[i][j] = 0
		*path = (*path)[:len(*path)-1]
	}

	// right
	if j+1 < n && (*visited)[i][j+1] != 1 && maze[i][j+1] == 1 {
		(*visited)[i][j] = 1
		*path = append(*path, 'R')
		solve(i, j+1, maze, n, ans, path, visited)
		(*visited)[i][j] = 0
		*path = (*path)[:len(*path)-1]
	}

	//upward
	if i-1 >= 0 && (*visited)[i-1][j] != 1 && maze[i-1][j] == 1 {
		(*visited)[i][j] = 1
		*path = append(*path, 'U')
		solve(i-1, j, maze, n, ans, path, visited)
		(*visited)[i][j] = 0
		*path = (*path)[:len(*path)-1]
	}
}
