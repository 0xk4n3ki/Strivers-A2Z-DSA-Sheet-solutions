package main

import "fmt"

func main() {
	arr := []int {2,2,1}
	num := xorApp(arr)
	fmt.Printf("%d appears only once in %v\n", num, arr)

	arr = []int {4, 1, 2, 1, 2}
	num = hashingUsingArr(arr)
	fmt.Printf("%d appears only once in %v\n", num, arr)

	arr = []int {4, 3, 67, 4, 3, 23, 23}
	num = hashingUsingMap(arr)
	fmt.Printf("%d appears only once in %v\n", num, arr)
}

func xorApp(arr []int) int {
	n := 0
	for _, i := range arr {
		n ^= i
	}
	return n
}

func hashingUsingArr(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if max < i {
			max = i
		}
	}

	arr1 := make([]int, max+1)
	for _, i := range arr {
		arr1[i]++
	}

	for i, val := range arr1 {
		if val == 1 {
			return i
		}
	}
	return -1
}

func hashingUsingMap(arr []int) int {
	map1 := map[int]int {}
	for _, i := range arr {
		map1[i]++
	}

	for i, val := range map1 {
		if val == 1 {
			return i
		}
	}
	return -1
}