package main

import "fmt"

func main() {
	arr := []int {10, 20, 30, 21, 23}
	ans := check(arr)
	fmt.Println("ans:", ans)

	arr = []int {10, 20, 30, 25, 15}
	ans = check(arr)
	fmt.Println("ans:", ans)
}

func check(arr []int) bool {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		if arr[i] > arr[2*i+1] || arr[i] > arr[2*i+2] {
			return false
		}
	}
	return true
}