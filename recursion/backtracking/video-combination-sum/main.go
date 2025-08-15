package main
import "fmt"

func main() {
	arr := []int {2, 3, 6, 7}
	target := 7
	ans := getComb(arr, target)
	fmt.Printf("ans: %v\n", ans)
}

func getComb(arr []int, target int) [][]int {
	ans := [][]int {}
	sub := []int {}


	combRec(arr, 0, target, &sub, &ans)
	return ans
}

func combRec(arr []int, i, target int, sub *[]int, ans *[][]int) {
	if i >= len(arr) {
		if target == 0 {
			tmp := []int {}
			tmp = append(tmp, *sub...)
			*ans = append(*ans, tmp)
		}
		return
	}

	if arr[i] <= target {
		*sub = append(*sub, arr[i])
		combRec(arr, i, target-arr[i], sub, ans)
		*sub = (*sub)[:len(arr)-1]
	}

	combRec(arr, i+1, target, sub, ans)
}