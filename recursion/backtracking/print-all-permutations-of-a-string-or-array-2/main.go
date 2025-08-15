package main
import "fmt"

func main() {
	arr := []int {1, 2, 3}
	ans := permFunc(arr)
	fmt.Printf("ans: %v\n", ans)
}

func permFunc(arr []int) [][]int {
	ans := [][]int {}

	permRec(arr, 0, &ans)
	return ans
}

func permRec(arr []int, i int, ans *[][]int) {
	if i == len(arr) {
		tmp := []int {}
		tmp = append(tmp, arr...)
		*ans = append(*ans, tmp)
		return
	}

	for j := i; j < len(arr); j++ {
		arr[j], arr[i] = arr[i], arr[j]
		permRec(arr, i+1, ans)
		arr[j], arr[i] = arr[i], arr[j]
	}
}