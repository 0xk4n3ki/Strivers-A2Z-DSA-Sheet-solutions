package main

import (
	"fmt"
)

func main() {
	arr1 := []int {1,2,3,4,5}  
	arr2 := []int {2,3,4,4,5}
	union := union1(arr1, arr2)
	fmt.Printf("arr1: %v, arr2: %v, union: %v\n", arr1, arr2, union)


	arr1 = []int {1,2,3,4,5,6,7,8,9,10}
	arr2 = []int {2,3,4,4,5,11,12}
	union = usingMap(arr1, arr2)
	fmt.Printf("arr1: %v, arr2: %v, union: %v\n", arr1, arr2, union)


	arr1 = []int {3, 23, 23, 45, 45, 372}
	arr2 = []int {11, 25, 45, 89, 888}
	union = usingSet(arr1, arr2)
	fmt.Printf("arr1: %v, arr2: %v, union: %v\n", arr1, arr2, union)


	arr1 = []int {1, 3, 6, 9, 23, 27, 45, 46, 333}
	arr2 = []int {21, 62, 99, 300, 1234}
	union = twoPointer(arr1, arr2)
	fmt.Printf("arr1: %v, arr2: %v, union: %v\n", arr1, arr2, union)
}

func union1(arr1, arr2 []int) []int {
	arr := []int {}
	

	for _, i := range arr1 {
		if !exists(arr, i) {
			arr = append(arr, i)
		}
	}
	for _, i := range arr2 {
		if !exists(arr, i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func usingMap(arr1, arr2 []int) []int {
	m := map[int]int {}
	arr := []int {}

	for _, i := range arr1 {
		m[i]++
	}
	for _, i := range arr2 {
		m[i]++
	}
	for i, _ := range m {
		arr = append(arr, i)
	}
	return arr
}


func usingSet(arr1, arr2 []int) []int {
	unique := make(map[int]struct{})

	for _, num := range arr1 {
		unique[num] = struct{}{}
	}
	for _, num := range arr2 {
		unique[num] = struct{}{}
	}
	// fmt.Printf("unique: %v\n", unique)

	var union []int
	for num := range unique {
		union = append(union, num)
	}
	return union
}


func twoPointer(arr1, arr2 []int) []int {
	i, j := 0, 0
	var union []int
	for i < len(arr1) && j < len(arr2) {
		var last int
		if len(union) > 0 {
			last = union[len(union)-1]
		}

		if arr1[i] < arr2[j] {
			if len(union) == 0 || arr1[i] != last {
				union = append(union, arr1[i])
			}
			i++
		}else if arr2[j] < arr1[i] {
			if len(union) == 0 || arr2[j] != last {
				union = append(union, arr2[j])
			}
			j++
		}else {
			if len(union) == 0 || arr1[i] != last {
				union = append(union, arr1[i])
			}
			i++
			j++
		}
	}

	for i < len(arr1) {
		if union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	for j < len(arr2) {
		if union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}
	return union
}




func exists(arr []int, n int) bool {
	for _, i := range arr {
		if i == n {
			return true
		}
	}
	return false
}