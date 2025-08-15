package main
import "fmt"

func main() {
	arr := []int {1, 2, 3, 2, 2}
	ans := brute(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	ans = optimal(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	ans = mostOptimal(arr)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {1, 2, 3, 2, 2}
	ans = mostOptimal(arr)
	fmt.Printf("ans: %d\n", ans)
}

func mostOptimal(arr []int) int {
	maxLen := 0
	numMap := map[int]int {}

	l, r := 0, 0
	for r < len(arr) {
		numMap[arr[r]]++

		if len(numMap) > 2 {
			numMap[arr[l]]--
			if numMap[arr[l]] == 0 {
				delete(numMap, arr[l])
			}
			l++
		}

		if len(numMap) <= 2 {
			n := r - l + 1
			if n > maxLen {
				maxLen = n
			}
		}
		r++
	}
	return maxLen
}

func optimal(arr []int) int {
	maxLen := 0
	numMap := map[int]int {}

	l, r := 0, 0
	for r < len(arr) {
		numMap[arr[r]]++

		for len(numMap) > 2 {
			numMap[arr[l]]--
			if numMap[arr[l]] == 0 {
				delete(numMap, arr[l])
			}
			l++
		}
		
		if len(numMap) <= 2 {
			n := r - l + 1
			if n > maxLen {
				maxLen = n
			}
		}
		r++
	}
	return maxLen
}

func brute(arr []int) int {
	maxlen := 0

	for i := 0; i < len(arr); i++ {
		setMap := map[int]struct{} {}
		for j := i; j < len(arr); j++ {
			setMap[arr[j]] = struct{}{}

			if len(setMap) <= 2 {
				n := j-i+1
				if n > maxlen {
					maxlen = n
				}
			}else {
				break
			}
		}
	}
	return maxlen
}