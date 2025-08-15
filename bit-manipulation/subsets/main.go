package main
import "fmt"

func main() {
	nums := []int {1, 2, 3}
	ans := subsets(nums)
	fmt.Printf("nums: %v, ans: %v\n", nums, ans)
}

func subsets(nums []int) [][]int {
	ans := [][]int {}
	n := len(nums)

	for i:= 0; i < (1<<n); i++ {
		tmp := []int {}
		for j := 0; j < n; j++ {
			if i & (1<<j) != 0 {
				tmp = append(tmp, nums[j])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}