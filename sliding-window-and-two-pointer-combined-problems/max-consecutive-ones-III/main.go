package main
import "fmt"

func main() {
	arr := []int {1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	ans := brute(arr, k)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
	k = 3
	ans = optimal(arr, k)
	fmt.Printf("ans: %d\n", ans)

	arr = []int {0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
	k = 3
	ans = mostOptimal(arr, k)
	fmt.Printf("ans: %d\n", ans)
}

func mostOptimal(arr []int, k int) int {
	maxlen := 0
	zeros := 0
	l, r := 0, 0

	for r < len(arr) {
		if arr[r] == 0 {
			zeros++
		}

		if zeros > k {
			if arr[l] == 0 {
				zeros--
			}
			l++
		}
		if zeros <= k {
			n := r - l + 1
			if n > maxlen {
				maxlen = n
			}
		}
		r++
	}
	return maxlen
}

func optimal(arr []int, k int) int {
	maxlen := 0
	l, r := 0, 0

	zeros := 0
	for r < len(arr) {
		if arr[r] == 0 {
			zeros++
		}

		for zeros > k {
			if arr[l] == 0 {
				zeros--
			}
			l++
		}

		n := r - l + 1
		if n > maxlen {
			maxlen = n
		}
		r++
	}
	return maxlen
}

func brute(arr []int, k int) int {
	maxlen := 0

	for i := 0; i < len(arr); i++ {
		zeros := 0
		for j := i; j < len(arr); j++ {
			if arr[j] == 0 {
				zeros++
			}

			if zeros <= k {
				n := j - i + 1
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