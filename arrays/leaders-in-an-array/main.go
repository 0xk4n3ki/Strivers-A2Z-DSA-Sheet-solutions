package main

import "fmt"

func main() {
	arr := []int {10, 22, 12, 3, 0, 6}
	basicApp(arr)
}

func basicApp(arr []int) {
	max := -1
	res := make([]int, 0)

	for i := len(arr)-1; i >= 0; i-- {
		if max < arr[i] {
			res = append(res, arr[i])
			max = arr[i]
		}
	}
	fmt.Printf("result: %v\n", res)
}