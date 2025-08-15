package main

import "fmt"

func main() {
	arr := []int {1,2,-4,-5}
	fmt.Printf("before: %v ", arr)
	basicApp(arr)
	fmt.Printf("after: %v\n", arr)


	arr = []int {1,2,-3,-1,-2,3}
	fmt.Printf("before: %v ", arr)
	optimalApp(arr)
	fmt.Printf("after: %v\n", arr)

	// for unequal number of +ve and -ve numbers
	arr = []int {1,2,-4,-5,3,4}
	fmt.Printf("before: %v ", arr)
	unequalApp(arr)
	fmt.Printf("after: %v\n", arr)
}

func basicApp(arr []int) {
	pos := []int {}
	neg := []int {}

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pos = append(pos, arr[i])
		}else {
			neg = append(neg, arr[i])
		}
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[2*i] = pos[i]
		arr[2*i+1] = neg[i]
	}
}

func optimalApp(arr []int) {
	posIndex, negIndex := 0, 1
	ans := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			ans[negIndex] = arr[i]
			negIndex += 2
		}else {
			ans[posIndex] = arr[i]
			posIndex += 2
		}
	}
}

func unequalApp(arr []int) {
	pos, neg := []int {}, []int {}
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pos = append(pos, arr[i])
		}else {
			neg = append(neg, arr[i])
		}
	}

	if len(pos) < len(neg) {
		for i := 0; i< len(pos); i++ {
			arr[2*i] = pos[i]
			arr[2*i+1] = neg[i]
		}

		index := len(pos) * 2
		for i := len(pos); i < len(neg); i++ {
			arr[index] = neg[i]
			index++
		}
	}else {
		for i := 0; i < len(neg); i++ {
			arr[2*i] = pos[i]
			arr[2*i+1] = neg[i]
		}

		index := len(neg) * 2
		for i := len(neg); i < len(pos); i++ {
			arr[index] = pos[i]
			index++
		}
	}
}