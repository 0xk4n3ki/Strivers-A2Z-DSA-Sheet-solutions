package main

import (
	"fmt"
)

func main() {
	image := [][]int {{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc, newColor := 1, 1, 2
	calc(image, sr, sc, newColor)
	fmt.Printf("ans: %v\n", image)

	image = [][]int {{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	dfs(image, sr, sc, newColor)
	fmt.Printf("ans: %v\n", image)

	image = [][]int {{1, 1, 1}, {2, 2, 0}, {2, 2, 2}}
	sr, sc, newColor = 2, 0, 3
	calc(image, sr, sc, newColor)
	fmt.Printf("ans: %v\n", image)

	image = [][]int {{1, 1, 1}, {2, 2, 0}, {2, 2, 2}}
	dfs(image, sr, sc, newColor)
	fmt.Printf("ans: %v\n", image)
}

func dfs(image [][]int, sr, sc, newColor int) {
	color := image[sr][sc]
	drow, dcol := []int {-1, 0, +1, 0}, []int {0, +1, 0, -1}

	recursion(image, sr, sc, newColor, color, drow, dcol)
}

func recursion(image [][]int, sr, sc, newColor, color int, drow, dcol []int) {
	image[sr][sc] = newColor
	n, m := len(image), len(image[0])

	for i := 0; i < 4; i++ {
		nrow, ncol := sr+drow[i], sc+dcol[i]
		if nrow >= 0 && nrow < n && ncol >= 0 && ncol < m && image[nrow][ncol] == color {
			recursion(image, nrow, ncol, newColor, color, drow, dcol)
		}
	}
}

func calc(image [][]int, sr, sc, newColor int) {
	rows, cols := len(image), len(image[0])
	vis := make([][]int, rows)
	for i := 0; i < rows; i++ {
		vis[i] = make([]int, cols)
	}

	q := new(Queue)
	q.Push(sr, sc)
	vis[sr][sc] = 1
	color := image[sr][sc]
	image[sr][sc] = newColor
	drow, dcol := []int {-1, 0, +1, 0}, []int {0, +1, 0, -1}
	for !q.IsEmpty() {
		row, col := q.Top().x, q.Top().y
		q.Pop()

		for i := 0; i < 4; i++ {
			x, y := row+drow[i], col+dcol[i]
			if x >= 0 && x < rows && y >= 0 && y < cols && image[x][y] == color && vis[x][y] != 1 {
				q.Push(x, y)
				vis[x][y] = 1
				image[x][y] = newColor
			}
		}
	}
}

type Pair struct {
	x, y int
}

type Queue struct {
	arr []Pair
}

func (q *Queue) Push(x, y int) {
	q.arr = append(q.arr, Pair{x, y})
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Top() Pair {
	if q.IsEmpty() {
		return Pair{}
	}
	return q.arr[0]
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[1:]
}