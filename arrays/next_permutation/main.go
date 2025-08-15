package main

import (
	"fmt"
)

func main() {
	arr := []int {1, 3, 2}
	permutations(arr)

	arr = []int {2, 1, 5, 4, 3, 0, 0}
	optimalApp(arr)
}

func optimalApp(arr []int) {
	var index int
	for index = len(arr)-2; index >= 0; index-- {
		if arr[index] < arr[index+1] {
			break
		}
	}
	if index == -1 {
		reverseArr(arr, 0, len(arr)-1)
		return
	}

	for i := len(arr)-1; i > index; i-- {
		if arr[i] > arr[index] {
			arr[i], arr[index] = arr[index], arr[i]
			break
		}
	}

	reverseArr(arr, index+1, len(arr)-1)
	
	fmt.Printf("resulting permutation: %v\n", arr)
}

func reverseArr(arr []int, low, high int) {
	for i, j := low, high; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func permutations(arr []int) {
	org := make([]int, len(arr))
	copy(org, arr)

	quickSort(arr, 0, len(arr)-1)

	ans := make([][]int, 0)
	recPermutations(arr, &ans, 0)
	// fmt.Printf("final ans: %v\n", ans)

	for index, val := range ans {
		i := 0
		for ; i < len(org); i++ {
			if val[i] != org[i] {
				break
			}
		}
		if i == len(org) {
			fmt.Printf("answer using permutations: %v\n", ans[index+1])
		}
	}
}

func recPermutations(arr []int, ans *[][]int, index int) {
	if index == len(arr) {
		perm := make([]int, len(arr))
		copy(perm, arr)
		*ans = append(*ans, perm)
		return
	}

	for i := index; i < len(arr); i++ {
		arr[i], arr[index] = arr[index], arr[i]
		recPermutations(arr, ans, index+1)
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		quickSort(arr, low, partitionIndex-1)
		quickSort(arr, partitionIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low, high
	for i < j {
		for i <= high - 1 && arr[i] <= pivot {
			i++
		}
		for j >= low + 1 && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[low] = arr[low], arr[j]
	return j
}