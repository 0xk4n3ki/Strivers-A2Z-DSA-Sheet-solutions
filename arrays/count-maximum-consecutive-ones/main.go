package main

import "fmt"

func main () {
	arr := []int {1, 1, 0, 1, 1, 1}
	n := basicApp(arr)
	fmt.Printf("count of maximum consecutive one in arr %v is: %d\n", arr, n)
}

func basicApp(arr []int) int {
	count := 0
	max := 0
	for _, i := range arr {
		if i == 1 {
			count++
			if count > max {
				max = count
			}
		}else {
			count = 0
		}
	}
	return max
}