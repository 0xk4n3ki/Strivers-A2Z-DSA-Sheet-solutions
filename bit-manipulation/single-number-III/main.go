package main
import "fmt"

func main() {
	nums := []int {1, 2, 1, 3, 5, 2}
	ans := single(nums)
	fmt.Printf("nums: %v, ans: %v\n", nums, ans)
}

func single(nums []int) []int {
	combXor := 0
	for _, i := range nums {
		combXor ^= i
	}

	lastBit := (combXor & (combXor-1)) ^ combXor

	setNum, unsetNum := 0, 0

	for _, i := range nums {
		if i & lastBit != 0 {
			setNum ^= i
		}else {
			unsetNum ^= i
		}
	}

	if setNum > unsetNum {
		return []int {unsetNum, setNum}
	}
	return []int {setNum, unsetNum}
}