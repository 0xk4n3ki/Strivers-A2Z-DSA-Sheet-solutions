package main

import (
	"fmt"
)

func main() {
	grid := [][]int {{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Printf("grid: %v, ans: %d\n", grid, calc(grid))

	grid = [][]int {{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Printf("grid: %v, ans: %d\n", grid, calc(grid))
}

func calc(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	q := new(Queue)
	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				q.Push(i, j, 0)
				vis[i][j] = 2
			}
		}
	}

	ans := 0
	drow, dcol := []int {-1, 0, +1, 0}, []int {0, +1, 0, -1}
	for !q.IsEmpty() {
		row, col := q.Top().pair.row, q.Top().pair.col
		t := q.Top().time
		if t > ans {
			ans = t
		}

		q.Pop()

		for i := 0; i < 4; i++ {
			r, c := row + drow[i], col + dcol[i]
			if (r >= 0 && r < n) && (c >= 0 && c < m) && grid[r][c] == 1 && vis[r][c] != 2 {
				q.Push(r, c, t+1)
				vis[r][c] = 2
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] != 2 && grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}

type Pair struct {
	row, col int
}

type Node struct {
	pair Pair
	time int
}

type Queue struct {
	arr []Node
}

func (q *Queue) Push(x, y, t int) {
	q.arr = append(q.arr, Node{Pair{x, y}, t})
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Top() Node {
	if q.IsEmpty() {
		return Node{}
	}
	return q.arr[0]
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[1:]
}