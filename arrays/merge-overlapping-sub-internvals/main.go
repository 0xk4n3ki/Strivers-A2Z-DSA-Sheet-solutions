package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int {{1, 3}, {2, 6}, {8, 9}, {9, 11}, {8, 10}, {2, 4}, {15, 18}, {16, 17}}

	bruteForceApp(arr)

	optimalApp(arr)
}


func optimalApp(arr [][]int) {
	sort.Slice(arr,func(i, j int) bool {
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})

	ans := [][]int {}
	
	for i := 0; i < len(arr); i++ {
		n := len(ans)

		if n == 0 || arr[i][0] > ans[n-1][1] {
			ans = append(ans, arr[i])
		}else {
			if ans[n-1][1] < arr[i][1] {
				ans[n-1][1] = arr[i][1]
			}
		}
	}

	fmt.Printf("ans: %v\n", ans)
}


func bruteForceApp(arr [][]int) {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})

	fmt.Printf("sorted intervals: %v\n", arr)

	ans := [][]int {}

	for i := 0; i < len(arr); i++ {
		start, end := arr[i][0], arr[i][1]

		if len(ans) != 0 && end <= ans[len(ans)-1][1] {
			continue
		}

		for j := i+1; j < len(arr); j++ {
			if arr[j][0] <= end {
				if end < arr[j][1] {
					end = arr[j][1]
				}else {
					break
				}
			}
		}
		ans = append(ans, []int {start, end})
	}

	fmt.Printf("ans: %v\n", ans)
}