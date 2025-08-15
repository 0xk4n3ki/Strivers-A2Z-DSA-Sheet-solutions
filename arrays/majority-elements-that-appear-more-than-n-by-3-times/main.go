package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int {1, 2, 2, 3, 2}

	usingMap(arr)

	arr = []int {11, 33, 33, 11, 33, 11}
	optimalApp(arr)
}


func optimalApp(arr []int) {
	count1, count2 := 0, 0
	ele1, ele2 := math.MinInt, math.MinInt

	for i := 0; i < len(arr); i++ {
		if count1 == 0 && ele1 != arr[i] {
			count1 = 1
			ele1 = arr[i]
		}else if count2 == 0 && ele2 != arr[i] {
			count2 = 1
			ele2 = arr[i]
		}else if arr[i] == ele1 {
			count1++
		}else if arr[i] == ele2 {
			count2++
		}else{
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, val := range arr {
		if val == ele1 {
			count1++
		}
		if val == ele2 {
			count2++
		}
	}

	ans := make([]int, 0)
	if count1 > int(len(arr)/3) {
		ans = append(ans, ele1)
	}
	if count2 > int(len(arr)/3) {
		ans = append(ans, ele2)
	}

	fmt.Printf("ans: %v\n", ans)
}


func usingMap(arr []int) {
	m := make(map[int]int)
	for _, val := range arr {
		m[val]++
	}

	count := 0
	ans := make([]int, 0)

	for _, c := range m {
		if c > len(arr)/3 {
			count++
			ans = append(ans, c)
		}
	}
	fmt.Printf("%v, count: %d, elements: %v\n", m, count, ans)
}