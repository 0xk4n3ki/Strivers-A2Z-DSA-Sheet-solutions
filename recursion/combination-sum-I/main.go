package main
import "fmt"

func main() {
	arr := []int {2, 3, 6, 7}
	target := 7

	ans := comb(arr, target)
	fmt.Printf("arr: %v, target: %d, ans: %v\n", arr, target, ans)
}

func comb(arr []int, target int) [][]int {
	ans := [][]int {}
	subseq := []int {}

	backTracking(0, arr, target, &subseq, &ans)

	return ans
}

func backTracking(index int, arr []int, target int, subseq *[]int, ans *[][]int) {
	if target == 0 {
		tmp := []int {}
		tmp = append(tmp, *subseq...)
		*ans = append(*ans, tmp)
		return
	}
	if index >= len(arr) {
		return
	}

	if target >= arr[index] {
		*subseq = append(*subseq, arr[index])
		backTracking(index, arr, target - arr[index], subseq, ans)
		*subseq = (*subseq)[:len(*subseq)-1]
	}

	backTracking(index+1, arr, target, subseq, ans)
}