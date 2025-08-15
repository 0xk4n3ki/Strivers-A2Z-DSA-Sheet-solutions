package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {1, 2, 3, 4, 5, 0}
	
	bruteForceApp(arr)
	betterApp(arr)
	kadaneAlgo(arr)
}


func betterApp(arr []int) {
	maxProduct := math.MinInt
	prefix, suffix := 1, 1

	for i, j := 0, len(arr)-1; i < len(arr) && j >= 0; i, j = i + 1, j - 1 {
		prefix *= arr[i]
		suffix *= arr[j]

		maxProduct = maxNum(maxProduct, prefix, suffix)

		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
	}
	
	fmt.Printf("arr: %v, max product: %d\n", arr, maxProduct)
}

func kadaneAlgo(arr []int) {
	maxProduct := arr[0]
	maxTillIndex, minTillIndex := arr[0], arr[0]

	for i := 0; i < len(arr); i++ {
		tmp := maxNum(arr[i], maxTillIndex * arr[i], minTillIndex * arr[i])
		minTillIndex = minNum(arr[i], maxProduct * arr[i], minTillIndex * arr[i])
		maxTillIndex = tmp

		if maxProduct < maxTillIndex {
			maxProduct = maxTillIndex
		}
	}

	fmt.Printf("arr: %v, max product: %d\n", arr, maxProduct)
}


func bruteForceApp(arr []int) {
	maxProduct := math.MinInt

	for i := 0; i < len(arr); i++ {
		product := 1
		for j := i; j < len(arr); j++ {
			product *= arr[j]
			if product > maxProduct {
				maxProduct = product
			}
		}
	}

	fmt.Printf("arr: %v, max product: %d\n", arr, maxProduct)
}


func maxNum(a, b, c int) int {
	if a > b && a > c {
		return a
	}else if b > a && b > c {
		return b
	}
	return c
}

func minNum(a, b, c int) int {
	if a < b && a < c {
		return a
	}else if b < c && b < a {
		return b
	}
	return c
}