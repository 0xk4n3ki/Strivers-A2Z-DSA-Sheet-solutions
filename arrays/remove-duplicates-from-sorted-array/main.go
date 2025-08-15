package main

import "fmt"

func main() {
	arr1 := []int {1,1,2,2,2,3,3}
	arr2 := []int {1,1,1,2,2,3,3,3,3,4,4}

	
	fmt.Printf("before: %v,", arr1) 
	n := remDups(arr1)
	fmt.Printf("after removing duplicates: %v\n", arr1[:n])
	
	fmt.Printf("before: %v,", arr2)
	n = remDups(arr2)
	fmt.Printf("after removing duplicates: %v\n", arr2[:n])
}

func remDups(arr []int) int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		} 
	}
	return j+1
}