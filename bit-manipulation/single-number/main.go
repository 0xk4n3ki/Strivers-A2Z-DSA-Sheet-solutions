package main
import "fmt"

func main() {
	arr := []int {2, 2, 1}
	ans := find(arr)
	fmt.Printf("arr: %v, ans: %d\n", arr, ans)

	arr = []int {4, 1, 2, 1, 2}
	ans = find(arr)
	fmt.Printf("arr: %v, ans: %d\n", arr, ans)
}

func find(arr []int) int {
	xor := 0
	for _, i := range arr {
		xor ^= i
	}
	return xor
}