package main
import "fmt"

func main() {
	fmt.Printf("n: %d, ans: %d\n", 10, calc(10))
	fmt.Printf("n: %d, ans: %d\n", 0, calc(0))
	fmt.Printf("n: %d, ans: %d\n", 1, calc(1))
}

func calc(n int) int {
	if n <= 2 {
		return 0
	}

	seArr := make([]bool, n)
	cnt := 1

	for i := 3; i < n; i += 2 {
		if !seArr[i] {
			cnt++
			for j := i*i; j < n; j += 2*i {
				seArr[j] = true
			}		
		}
	}

	return cnt
}