package main
import "fmt"

func main() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	ans := findPath2(maze, len(maze))
	fmt.Printf("ans: %v\n", ans)
}

func findPath2(maze [][]int, n int) []string {
	ans := []string {}
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}

	path := []rune {}

	// DLRU
	di := []int {1, 0, 0, -1}
	dj := []int {0, -1, 1, 0}

	if maze[0][0] == 1 {
		solve2(0, 0, maze, n, &ans, &path, &visited, di, dj)
	}
	return ans
}

func solve2(i, j int, maze [][]int, n int, ans *[]string, path *[]rune, visited *[][]int, di, dj []int) {
	if i == n-1 && j == n-1 {
		*ans = append(*ans, string(*path))
		return
	}

	direction := "DLRU"

	for ind := 0; ind < 4; ind++ {
		nexti := i + di[ind]
		nextj := j + dj[ind]

		if nexti >= 0 && nexti < n &&
		nextj >= 0 && nextj < n &&
		(*visited)[nexti][nextj] != 1 && maze[nexti][nextj] == 1 {
			(*visited)[i][j] = 1
			*path = append(*path, rune(direction[ind]))
			solve2(nexti, nextj, maze, n, ans, path, visited, di, dj)
			(*visited)[i][j] = 0
			*path = (*path)[:len(*path)-1]
		}
	}
}