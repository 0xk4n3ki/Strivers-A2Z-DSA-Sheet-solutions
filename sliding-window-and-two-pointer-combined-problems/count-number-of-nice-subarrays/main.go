package main
import "fmt"

func main() {
	arr := []int {1, 1, 2, 1, 1}
	k := 3
	ans := niceNum(arr, k)-niceNum(arr, k-1)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {2, 4, 6}
	k = 1
	ans = niceNum(arr, k)-niceNum(arr, k-1)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {2,2,2,1,2,2,1,2,2,2}
	k = 2
	ans = niceNum(arr, k)-niceNum(arr, k-1)
	fmt.Printf("ans: %d\n", ans)
}

func niceNum(arr []int, k int) int {
	if k < 0 {
		return 0
	}
	ans := 0
	oddCnt := 0
	i, j := 0, 0

	for j < len(arr) {
		oddCnt += (arr[j] % 2)

		for oddCnt > k {
			fmt.Printf("arr: %v, i: %d, j: %d, oddCnt: %d, k: %d\n", arr, i, j, oddCnt, k)
			oddCnt -= (arr[i] % 2)
			i++
		}
		ans += j-i+1
		j++
	}
	return ans
}