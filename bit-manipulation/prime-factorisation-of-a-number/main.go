package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {2, 3, 4, 5, 6, 18}
	ans := getPrime(arr)
	fmt.Printf("ans: %v\n", ans)
}

func getPrime(arr []int) [][]int {
	max := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	spf := make([]int, max+1)
	for i := 2; i <= max; i++ {
		if spf[i] == 0 {
			for j := i; j <= max; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}
	fmt.Printf("spf: %v\n", spf)

	res := [][]int {}
	for i := 0; i < len(arr); i++ {
		tmp := []int {}

		for arr[i] > 1 {
			tmp = append(tmp, spf[arr[i]])
			arr[i] /= spf[arr[i]]
		}

		res = append(res, tmp)
	}

	return res
}