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
	fmt.Print("Enter nnumber of nodes(n) and edges(m): ")
	fmt.Scan(&n, &m)

	adjList := make([][]int, n+1)

	for i := 1; i <= m; i++ {
		var j, k int
		fmt.Printf("[%d]: ", i)
		fmt.Scan(&j, &k)

		adjList[j] = append(adjList[j], k)
		adjList[k] = append(adjList[k], j)
	}

	ans := []int {}
	vis := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if vis[i] != 1 {
			ans = append(ans, bfs(adjList, &vis, i)...)
		}
	}

	fmt.Printf("ans: %v\n", ans)
}

func bfs(adjList [][]int, vis *[]int, t int) []int {
	q := new(Queue)
	ans := []int {}

	q.Push(t)
	(*vis)[t] = 1

	for !q.IsEmpty() {
		tmp := q.Top()
		ans = append(ans, tmp)
		q.Pop()

		for _, i := range adjList[tmp] {
			if (*vis)[i] != 1 {
				q.Push(i)
				(*vis)[i] = 1
			}
		}
	}
	return ans
}


type Queue struct {
	Arr []int
}

func (q *Queue) Push(data int) {
	q.Arr = append(q.Arr, data)
}

func (q *Queue) Size() int {
	return len(q.Arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Top() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Arr[0]
}

func (q *Queue) Pop() {
	if q.IsEmpty() {
		return
	}
	q.Arr = q.Arr[1:]
}