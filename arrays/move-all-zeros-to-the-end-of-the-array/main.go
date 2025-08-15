package main

import "fmt"

func main() {
	arr1 := []int {1 ,0 ,2 ,3 ,0 ,4 ,0 ,1}
	arr2 := []int {1,2,0,1,0,4,0}

	fmt.Printf("before: %v ", arr1)
	otherApp(arr1)
	fmt.Printf("after moving: %v\n", arr1)

	fmt.Printf("before: %v ", arr2)
	optimalApp(arr2)
	fmt.Printf("after moving: %v\n", arr2)
}

func optimalApp(arr []int) {
	i, j, ele := 0, 0, 0
	for i, ele = range arr {
		if ele != 0 {
			arr[j] = arr[i]
			j++
		}
	}
	for ; j <= i; j++ {
		arr[j] = 0
	}	
}

func otherApp(arr []int) {
	j := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}
	if j == -1 {
		return
	}

	for i := j+1; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
}