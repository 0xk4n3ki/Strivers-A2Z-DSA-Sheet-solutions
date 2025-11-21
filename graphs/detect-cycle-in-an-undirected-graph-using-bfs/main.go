package main

import (
	"fmt"
)

func main() {
	v, e := 8, 7
	edges := [][]int {{1, 2}, {1, 3}, {2, 5}, {2, 6}, {3, 4}, {3, 7}, {7, 8}}
	fmt.Printf("ans: %t\n", check(edges, v, e))

	v, e = 8, 6
	edges = [][]int {{1, 2}, {2, 3}, {4, 5}, {5, 6}, {4, 6}, {7, 8}}
	fmt.Printf("ans: %t\n", check(edges, v, e))

	v, e = 7, 7
	edges = [][]int {{1, 2}, {1, 3}, {3, 6}, {3, 4}, {6, 7}, {2, 5}, {5, 7}}
	fmt.Printf("ans: %t\n", check(edges, v, e))
}

func check(edges [][]int, v, e int) bool {
	vis := make([]int, v+1)
	adj := make([][]int, v+1)
	for i := 0; i < e; i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	fmt.Printf("adj: %v\n", adj)

	for i := 1; i <= v; i++ {
		if vis[i] != 1 {
			if bfs(adj, &vis, i) {
				return true
			}
		}
	}
	return false
}

func bfs(adj [][]int, vis *[]int, i int) bool {
	q := new(Queue)
	q.Push(i, -1)
	(*vis)[i] = 1

	for !q.IsEmpty() {
		s, p := q.Top()
		q.Pop()

		for _, j := range adj[s] {
			if (*vis)[j] == 1 && j != p {
				return true
			}else if (*vis)[j] != 1 {
				q.Push(j, s)
				(*vis)[j] = 1
			}
		}
	}
	return false
}

type Pair struct {
	source, parent int
}

type Queue struct {
	arr []Pair
}

func (q * Queue) Push(source, parent int) {
	q.arr = append(q.arr, Pair{source, parent})
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q * Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Top() (int, int) {
	if q.IsEmpty() {
		return -1, -1
	}
	return q.arr[0].source, q.arr[0].parent
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[1:]
}