package main

import (
	"fmt"
)

func main() {
	grid := [][]int {{1, 0, 1}, {1, 1, 0}, {1, 0, 0}}
	fmt.Printf("ans: %v\n", calc(grid))

	grid = [][]int {{0, 0, 0}, {0, 1, 0}, {1, 0, 1}}
	fmt.Printf("ans: %v\n", calc(grid))
}

func calc(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	vis := make([][]int, n)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
		dist[i] = make([]int, m)
	}

	q := new(Queue)
	
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				q.Push(i, j, 0)
				vis[i][j] = 1
			}
		}
	}

	drow, dcol := []int {-1, 0, +1, 0}, []int {0, +1, 0, -1}
	for !q.IsEmpty() {
		row, col, steps := q.Top()
		q.Pop()

		dist[row][col] = steps

		for i := 0; i < 4; i++ {
			nrow, ncol := row + drow[i], col + dcol[i]
			if nrow >= 0 && ncol >= 0 && nrow < n && ncol < m && vis[nrow][ncol] != 1 {
				vis[nrow][ncol] = 1
				q.Push(nrow, ncol, steps+1)
			}
		}
	}
	return dist
}

type Node struct {
	x, y, step int
}

type Queue struct {
	arr []Node
}

func (q *Queue) Push(x, y, step int) {
	q.arr = append(q.arr, Node{x, y, step})
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Top() (int, int, int) {
	if q.IsEmpty() {
		return -1, -1, -1
	}
	return q.arr[0].x, q.arr[0].y, q.arr[0].step
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[1:]
}