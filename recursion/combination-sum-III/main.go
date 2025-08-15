package main
import "fmt"

func main() {
	n, k := 9, 3

	ans := comb(n, k)
	fmt.Printf("n: %d, k: %d, ans: %v\n", n, k, ans)
}

func comb(sum, num int) [][]int {
	ans := [][]int {}
	sub := []int {}

	backtracking(1, num, sum, &sub, &ans)

	return ans
}

func backtracking(start, num, sum int, sub *[]int, ans *[][]int) {
	if num == 0 {
		if sum == 0 {
			*ans = append(*ans, append([]int {}, *sub...))
		}
		return
	}

	for i := start; i <= 9; i++ {
		if sum < i {
			break
		}
		*sub = append(*sub, i)
		backtracking(i+1, num-1, sum-i, sub, ans)
		*sub = (*sub)[:len(*sub)-1]
	}
}