package main

import "fmt"


// test case:
// 10 8
// 1 2
// 1 3
// 3 4
// 2 4
// 5 6
// 5 7
// 6 7
// 8 9

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fmt.Printf("adj: %v\n", adj)
	ans := []int {}
	vis := make([]int, n+1)
	for i := 1; i <= n; i++ { // for the connected components 
		if vis[i] != 1 {
			vis[i] = 1
			ans = append(ans, bfs(adj, vis, i)...)
			fmt.Printf("vis: %v\n", vis)
		}
	}

	fmt.Printf("ans: %v\n", ans)
}

func bfs(adj [][]int, vis []int, start int) []int {
	q := new(Queue)
	q.Push(start)
	bfs := []int {}

	for !q.IsEmpty() {
		node := q.Top()
		q.Pop()
		bfs = append(bfs, node)

		for _, i := range adj[node] {
			if vis[i] != 1 {
				vis[i] = 1
				q.Push(i)
			}
		}
	}
	return bfs
}

type Queue struct {
	arr []int
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Push(val int) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Top() int {
	if q.IsEmpty() {
		return -1
	}
	return q.arr[q.Size()-1]
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.arr = q.arr[:q.Size()-1]
}