package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {12, 34, 67, 90}
	students := 2

	bs(arr, students)

	arr = []int {25, 46, 28, 49, 24}
	students = 4
	bs(arr, students)
}

func bs(arr []int, studentsLimit int) {
	if studentsLimit > len(arr) {
		fmt.Printf("arr: %v, students: %d, pages: %d\n", arr, studentsLimit, -1)
	}

	low, high := maxEle(arr)

	for low <= high {
		mid := (low + high)/2

		students := countStudents(arr, mid)

		if students > studentsLimit {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	fmt.Printf("arr: %v, students: %d, pages: %d\n", arr, studentsLimit, low)
}

func countStudents(arr []int, pages int) int {
	cnt := 1
	sum := 0
	
	for i := 0; i < len(arr); i++ {
		if arr[i] + sum <= pages {
			sum += arr[i]
		} else {
			cnt++
			sum = arr[i]
		}
	}
	return cnt
}


func maxEle(arr []int) (int, int) {
	res := math.MinInt
	sum := 0
	for _, val := range arr {
		sum += val
		if res < val {
			res = val
		}
	}
	return res, sum
}