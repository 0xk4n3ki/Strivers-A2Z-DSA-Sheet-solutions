package main
import "fmt"

func main() {
	arr := []int {1, 0, 1, 0, 1}
	goal := 2
	ans := binarySum(arr, goal)-binarySum(arr, goal-1)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {0, 0, 0, 0, 0}
	goal = 0
	ans = binarySum(arr, goal)-binarySum(arr, goal-1)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {1, 0, 1, 0, 1}
	goal = 2
	ans = cntSub(arr, goal)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {0, 0, 0, 0, 0}
	goal = 0
	ans = cntSub(arr, goal)
	fmt.Printf("ans: %d\n", ans)
}

func binarySum(arr []int, goal int) int {
	if goal < 0 {
		return 0
	}
	cnt := 0
	sum := 0
	i, j := 0, 0
	n := len(arr)

	for j < n {
		sum += arr[j]

		for sum > goal {
			sum -= arr[i]
			i++
		}

		cnt += (j-i+1)
		j++  
	}
	return cnt
}


func cntSub(arr []int, goal int) int {
	hashMap := map[int]int {}
	cnt := 0
	sum := 0
	hashMap[0] = 1

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		rem := sum - goal
		if _, ok := hashMap[rem]; ok {
			cnt += hashMap[rem]
		}

		hashMap[sum]++
	}
	return cnt
}