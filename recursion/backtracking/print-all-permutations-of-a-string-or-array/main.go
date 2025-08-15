package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	ans := perm(arr)
	fmt.Printf("ans: %v\n", ans)
}

// approach 1: using a map to keep the track of used number
func perm(arr []int) [][]int {
	ans := [][]int{}
	sub := []int{}
	numMap := map[int]bool{}

	for i := 0; i < len(arr); i++ {
		numMap[arr[i]] = false
	}

	permRec(arr, &sub, &ans, &numMap)
	return ans
}

func permRec(arr []int, sub *[]int, ans *[][]int, numMap *map[int]bool) {
	if len(*sub) == len(arr) {
		tmp := []int{}
		tmp = append(tmp, *sub...)
		*ans = append(*ans, tmp)
		return
	}

	for j := 0; j < len(arr); j++ {
		if (*numMap)[arr[j]] {
			continue
		}
		*sub = append(*sub, arr[j])
		(*numMap)[arr[j]] = true
		permRec(arr, sub, ans, numMap)
		*sub = (*sub)[:len(*sub)-1]
		(*numMap)[arr[j]] = false
	}
}
