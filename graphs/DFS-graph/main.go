package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Print("Enter the number of nodes(n) and edges(m): ")
	fmt.Scan(&n, &m)

	adjList := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		fmt.Printf("[%d]: ", i)
		var u, v int
		fmt.Scan(&u, &v)

		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	ans := []int {}
	ans2 := []int {}
	vis := make([]int, n+1)
	vis2 := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if vis[i] != 1 {
			ans = append(ans, dfs(adjList, &vis, i)...)
		}
		if vis2[i] != 1 {
			recDFS(adjList, &ans2, &vis2, i)
		}
	}

	fmt.Printf("ans: %v\n", ans)
	fmt.Printf("ans2: %v\n", ans2)
}

func recDFS(adjList [][]int, ans, vis *[]int, i int) {
	*ans = append(*ans, i)
	(*vis)[i] = 1

	for _, j := range adjList[i] {
		if (*vis)[j] != 1 {
			recDFS(adjList, ans, vis, j)
		}
	}
}



func dfs(adjList [][]int, vis *[]int, t int) []int {
	st := new(Stack)
	ans := []int {}

	st.Push(t)
	(*vis)[t] = 1

	for !st.IsEmpty() {
		tmp := st.Top()
		st.Pop()
		ans = append(ans, tmp)

		for _, i := range adjList[tmp] {
			if (*vis)[i] != 1 {
				st.Push(i)
				(*vis)[i] = 1
			}
		}
	}
	return ans
}


type Stack struct {
	arr []int
}

func (st *Stack) Push(data int) {
	st.arr = append(st.arr, data)
}

func (st *Stack) Size() int {
	return len(st.arr)
}

func (st *Stack) IsEmpty() bool {
	return st.Size() == 0
}

func (st *Stack) Top() int {
	if st.IsEmpty() {
		return -1
	}
	return st.arr[st.Size()-1]
}

func (st *Stack) Pop() {
	if st.IsEmpty() {
		return
	}
	st.arr = st.arr[:st.Size()-1]
}