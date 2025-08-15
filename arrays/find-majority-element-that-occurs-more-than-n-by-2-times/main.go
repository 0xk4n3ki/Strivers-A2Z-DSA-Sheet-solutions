package main

import "fmt"

func main() {
	arr := []int{3, 2, 3}
	ele := basicApp(arr)
	fmt.Printf("%d occurs more then n/2 times in %v\n", ele, arr)

	arr = []int{2, 2, 1, 1, 1, 2, 2}
	ele = mooreApp(arr)
	fmt.Printf("%d occurs more then n/2 times in %v\n", ele, arr)
}

func basicApp(arr []int) int {
	count := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for i, ele := range count {
		if ele > int(len(arr)/2) {
			return i
		}
	}
	return -1
}

// Moore's Voting Algorithm
func mooreApp(arr []int) int {
	count := 0
	var ele int
	for i := 0; i < len(arr); i++ {
		if count == 0 {
			ele = arr[i]
			count++
		} else if ele == arr[i] {
			count++
		} else {
			count--
		}
	}

	count1 := 0
	for _, i := range arr {
		if i == ele {
			count1++
		}
	}
	if count1 > int(len(arr)/2) {
		return ele
	}
	return -1
}
