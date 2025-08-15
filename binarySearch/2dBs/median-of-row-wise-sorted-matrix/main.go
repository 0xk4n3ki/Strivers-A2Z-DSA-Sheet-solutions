package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int {{1, 4, 9}, {2, 5, 6}, {3, 8, 7}}

	bruteforce(matrix)
	optimalApp(matrix)
}

func optimalApp(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	low, high := math.MaxInt, math.MinInt

	for i := 0; i < n; i++ {
		if low > matrix[i][0] {
			low = matrix[i][0]
		}

		if high < matrix[i][m-1] {
			high = matrix[i][m-1]
		}
	}

	req := (n * m)/2

	for low <= high {
		mid := (low + high)/2

		smallEqual := countSmallEqual(matrix, mid)

		if smallEqual <= req {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}

	fmt.Printf("matrix: %v, median: %d\n", matrix, low)
}

func countSmallEqual(matrix [][]int, mid int) int {
	count := 0

	for i := 0; i < len(matrix); i++ {
		count += upperBound(matrix[i], mid)
	}
	return count
}

func upperBound(arr []int, x int) int {
	low, high := 0, len(arr)

	ans := len(arr)

	for low <= high {
		mid := (low + high)/2

		if arr[mid] > x {
			ans = mid
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return ans
}

func bruteforce(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	arr := []int {}

	for i := 0; i < n; i++ {
		arr = append(arr, matrix[i]...)
	}

	insertionSort(arr)
	
	ans := arr[(m*n)/2]

	fmt.Printf("matrix: %v, ans: %d\n", matrix, ans)
}

func insertionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}