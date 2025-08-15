package main

import "fmt"

func main() {
	arr := []int{6, -2, 2, -8, 1, 7, 4, -10}

	bruteforceApp(arr)
	betterApp(arr)
}

// maintain an map with thee sum till the index "map[sum] = index"
// now whenever a sum appears that is already present,
// that means uske baad abhi tak ka sum 0 hai
func betterApp(arr []int) {
	maxLen := 0
	sumMap := map[int]int{}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if sum == 0 {
			maxLen = i + 1
		} else if _, ok := sumMap[sum]; ok {
			len := i - sumMap[sum]
			if len > maxLen {
				maxLen = len
			}
		} else {
			sumMap[sum] = i
		}
	}
	fmt.Printf("maxLen: %d\n", maxLen)
}

func bruteforceApp(arr []int) {
	maxLen := -1
	maxI, maxJ := -1, -1

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]

			if sum == 0 {
				len := j - i + 1
				if len > maxLen {
					maxLen = len
					maxI, maxJ = i, j
				}
			}
		}
	}
	if maxLen == -1 {
		fmt.Printf("no such sub arrays\n")
		return
	}
	fmt.Printf("longest len: %d, subarray: %v\n", maxLen, arr[maxI:maxJ+1])
}
