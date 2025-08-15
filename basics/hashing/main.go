package main

import "fmt"

func main() {
	arr := []int {1, 2, 3, 1, 3, 2, 12, 2, 3, 2}
	
	mp := map[int] int {}
	for _, i := range arr {
		mp[i]++
	}

	query := []int {1, 2, 3, 4, 12}
	for _, i := range query {
		fmt.Printf("%d appeared %d times\n", i, mp[i])
	}

	maxFreq, minFreq := 0, len(arr)
	maxEle, minEle := 0, 0
	for key, value := range mp {
		if value > maxFreq {
			maxEle = key
			maxFreq = value
		}
		if value < minFreq {
			minEle = key
			minFreq = value
		}
	}
	fmt.Printf("%d appeared max times: %d\n", maxEle, maxFreq)
	fmt.Printf("%d appeared min times: %d\n", minEle, minFreq)
}